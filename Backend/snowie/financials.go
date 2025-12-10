package snowie

import (
	stripeclient "backend/clients"
	"backend/utils"
	"log"
	"net/http"
	"os"

	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FinancialStats struct {
	Revenue      int64  `json:"revenue"`
	Refunded     int64  `json:"refunded"`
	DisputesLost int64  `json:"disputes_lost"`
	Profit       int64  `json:"profit"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

func GetFinancialStats(c *gin.Context) {

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate, endDate = utils.ApplyDefaultMonth(startDate, endDate)
	}

	cacheKey := "snowie_financials:" + startDate + ":" + endDate
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data stripeclient.IAResponse
		if err := json.Unmarshal([]byte(cached), &data); err == nil {
			log.Printf("[Snowie] Financials cache HIT (%s)", cacheKey)
			utils.CustomResponse(c, http.StatusOK, true, "Invoice-based financial stats (cached)", data)
			return
		}
	}
	log.Printf("[Snowie] Financials cache MISS (%s)", cacheKey)
	log.Printf("[Snowie] IA request | %s → %s", startDate, endDate)

	keys := []string{
		os.Getenv("SNOWIE_STRIPE_KEY_1"),
		os.Getenv("SNOWIE_STRIPE_KEY_2"),
	}

	type msg struct {
		ia  *stripeclient.IAResult
		err error
	}

	ch := make(chan msg, len(keys))

	// -------- parallel fetch --------
	for _, key := range keys {
		go func(secret string) {
			client := stripeclient.New(secret)
			ia, err := client.GetInvoiceBasedIA(startDate, endDate)
			ch <- msg{ia: ia, err: err}
		}(key)
	}

	// -------- merged result --------
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

	// -------- collect --------
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

	// 1. Try cache
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data []stripeclient.MonthlyStats
		json.Unmarshal([]byte(cached), &data)
		utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved (cache)", data)
		return
	}

	keys := []string{
		os.Getenv("SNOWIE_STRIPE_KEY_1"),
		os.Getenv("SNOWIE_STRIPE_KEY_2"),
	}

	type result struct {
		stats []stripeclient.MonthlyStats
		err   error
	}

	statsCh := make(chan result, len(keys))

	// 2. Fetch both Stripe accounts in parallel
	for _, key := range keys {
		go func(secret string) {
			client := stripeclient.New(secret)
			data, e := client.GetMonthlyStats(year)
			statsCh <- result{stats: data, err: e}
		}(key)
	}

	// 3. Merge data
	var combined [12]stripeclient.MonthlyStats
	first := true

	for range keys {
		res := <-statsCh
		if res.err != nil {
			utils.CustomResponse(c, http.StatusInternalServerError, false, res.err.Error(), nil)
			return
		}

		// Initialize months on first pass
		if first {
			for i := 0; i < 12; i++ {
				combined[i] = res.stats[i]
			}
			first = false
			continue
		}

		// Merge subsequent accounts
		for i := 0; i < 12; i++ {
			combined[i].Revenue += res.stats[i].Revenue
			combined[i].Profit += res.stats[i].Profit
		}
	}

	// Convert fixed array to slice
	final := combined[:]

	// 4. Cache for 30 minutes (graphs don’t need real-time updates)
	b, _ := json.Marshal(final)
	utils.Set(cacheKey, string(b), 30*time.Minute)

	// 5. Return response
	utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved successfully", final)
}
