package closerx

import (
	stripeclient "backend/clients"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetFinancialStats(c *gin.Context) {

	start := time.Now()

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate, endDate = utils.ApplyDefaultMonth(startDate, endDate)
	}

	// -------------------------------
	// ✅ CACHE KEY
	// -------------------------------
	cacheKey := "closerx_financials:" + startDate + ":" + endDate

	// -------------------------------
	// ✅ CACHE HIT
	// -------------------------------
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data stripeclient.IAResponse
		if err := json.Unmarshal([]byte(cached), &data); err == nil {
			log.Printf("[CloserX] Financials cache HIT (%s)", cacheKey)

			utils.CustomResponse(
				c,
				http.StatusOK,
				true,
				"Invoice-based financial stats (cached)",
				data,
			)
			return
		}
	}

	log.Printf("[CloserX] Financials cache MISS (%s)", cacheKey)

	keys := []string{
		os.Getenv("CLOSERX_STRIPE_KEY_1"),
		os.Getenv("CLOSERX_STRIPE_KEY_2"),
	}

	type resultMsg struct {
		ia  *stripeclient.IAResult
		err error
	}

	ch := make(chan resultMsg, len(keys))

	// -------------------------------------------------
	// PARALLEL FETCH (PER ACCOUNT)
	// -------------------------------------------------
	for _, key := range keys {
		go func(secret string) {
			client := stripeclient.New(secret)
			ia, err := client.GetInvoiceBasedIA(startDate, endDate)
			ch <- resultMsg{ia: ia, err: err}
		}(key)
	}

	// -------------------------------------------------
	// FINAL MERGED RESULT
	// -------------------------------------------------
	final := &stripeclient.IAResult{
		Subscriptions: map[string]*stripeclient.Bucket{},
		Credits:       map[string]*stripeclient.Bucket{},
		Others:        map[string]*stripeclient.Bucket{},
	}

	merge := func(dst, src map[string]*stripeclient.Bucket) {
		for price, b := range src {
			if _, ok := dst[price]; !ok {
				dst[price] = &stripeclient.Bucket{}
			}
			dst[price].Count += b.Count
			dst[price].Revenue += b.Revenue
			dst[price].Refunded += b.Refunded
			dst[price].Disputes += b.Disputes
			dst[price].Profit += b.Profit
		}
	}

	// -------------------------------------------------
	// COLLECT
	// -------------------------------------------------
	for i := 0; i < len(keys); i++ {
		res := <-ch
		if res.err != nil {
			utils.CustomResponse(c, http.StatusInternalServerError, false, res.err.Error(), nil)
			return
		}

		merge(final.Subscriptions, res.ia.Subscriptions)
		merge(final.Credits, res.ia.Credits)
		merge(final.Others, res.ia.Others)
	}

	log.Printf("[CloserX] IA merged in %s", time.Since(start))

	formatted := stripeclient.FormatIAResponse(final)

	// -------------------------------
	// ✅ CACHE SAVE (5 minutes)
	// -------------------------------
	if b, err := json.Marshal(formatted); err == nil {
		_ = utils.Set(cacheKey, b, 5*time.Minute)
	}

	utils.CustomResponse(
		c,
		http.StatusOK,
		true,
		"Invoice-based financial stats (price-level merged)",
		formatted,
	)
}

// -----------------------
// MONTHLY STATS (FAST)
// -----------------------
func GetMonthlyStats(c *gin.Context) {
	yearStr := c.Query("year")
	if yearStr == "" {
		yearStr = "2025"
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		utils.CustomResponse(c, http.StatusBadRequest, false, "invalid year", nil)
		return
	}

	cacheKey := "closerx_monthly:" + yearStr

	// Cache hit?
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data []stripeclient.MonthlyStats
		json.Unmarshal([]byte(cached), &data)
		utils.CustomResponse(c, http.StatusOK, true, "Monthly stats (cache)", data)
		return
	}

	keys := []string{
		os.Getenv("CLOSERX_STRIPE_KEY_1"),
		os.Getenv("CLOSERX_STRIPE_KEY_2"),
	}

	type monthlyResult struct {
		stats []stripeclient.MonthlyStats
		err   error
	}

	resultCh := make(chan monthlyResult, len(keys))

	for _, key := range keys {
		go func(k string) {
			client := stripeclient.New(k)
			stats, e := client.GetMonthlyStats(year)
			resultCh <- monthlyResult{stats: stats, err: e}
		}(key)
	}

	// Merge results
	var combined [12]stripeclient.MonthlyStats
	first := true

	for range keys {
		res := <-resultCh
		if res.err != nil {
			utils.CustomResponse(c, http.StatusInternalServerError, false, res.err.Error(), nil)
			return
		}

		if first {
			copy(combined[:], res.stats)
			first = false
			continue
		}

		for i := 0; i < 12; i++ {
			combined[i].Revenue += res.stats[i].Revenue
			combined[i].Profit += res.stats[i].Profit
		}
	}

	final := combined[:]

	// Cache for 30 minutes
	b, _ := json.Marshal(final)
	utils.Set(cacheKey, string(b), 30*time.Minute)

	utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved successfully", final)
}
