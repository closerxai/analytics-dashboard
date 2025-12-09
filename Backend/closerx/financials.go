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

type FinancialStats struct {
	Revenue      int64  `json:"revenue"`
	Refunded     int64  `json:"refunded"`
	DisputesLost int64  `json:"disputes_lost"`
	Profit       int64  `json:"profit"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

// Fetch totals for one Stripe account using Balance Transactions
func fetchForAccount(secret, start, end string, ch chan<- FinancialStats, errCh chan<- error) {
	client := stripeclient.New(secret)

	r, f, d, err := client.GetTotals(start, end)
	if err != nil {
		errCh <- err
		return
	}

	ch <- FinancialStats{
		Revenue:      r,
		Refunded:     f,
		DisputesLost: d,
		Profit:       r - f - d,
		StartDate:    start,
		EndDate:      end,
	}
}

func GetFinancialStats(c *gin.Context) {
	reqStart := time.Now()

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate, endDate = utils.ApplyDefaultMonth(startDate, endDate)
	}

	log.Printf("[CloserX] Request | start=%s end=%s", startDate, endDate)

	cacheKey := "closerx:" + startDate + ":" + endDate

	// Cache check
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		log.Printf("[CloserX] Cache HIT → %s", cacheKey)
		var data FinancialStats
		json.Unmarshal([]byte(cached), &data)
		utils.CustomResponse(c, http.StatusOK, true, "Financial stats (cache)", data)
		return
	}

	log.Printf("[CloserX] Cache MISS → %s", cacheKey)

	keys := []string{
		os.Getenv("CLOSERX_STRIPE_KEY_1"),
		os.Getenv("CLOSERX_STRIPE_KEY_2"),
	}

	statsCh := make(chan FinancialStats, len(keys))
	errCh := make(chan error, len(keys))

	log.Printf("[CloserX] Fetching Stripe data for %d accounts", len(keys))

	for _, key := range keys {
		go func(k string) {
			safe := k[len(k)-4:]
			log.Printf("[CloserX] Fetching account ...%s", safe)
			fetchForAccount(k, startDate, endDate, statsCh, errCh)
		}(key)
	}

	var total FinancialStats

	for range keys {
		select {
		case s := <-statsCh:
			total.Revenue += s.Revenue
			total.Refunded += s.Refunded
			total.DisputesLost += s.DisputesLost
		case err := <-errCh:
			log.Printf("[CloserX] ERROR: %v", err)
			utils.CustomResponse(c, http.StatusInternalServerError, false, err.Error(), nil)
			return
		}
	}

	refAbs := utils.Abs64(total.Refunded)
	dispAbs := utils.Abs64(total.DisputesLost)

	total.Profit = total.Revenue - refAbs - dispAbs
	total.Refunded = refAbs
	total.DisputesLost = dispAbs
	total.StartDate = startDate
	total.EndDate = endDate

	log.Printf("[CloserX] Totals | revenue=%d refunded=%d disputes=%d profit=%d",
		total.Revenue, total.Refunded, total.DisputesLost, total.Profit)

	// Save to cache
	b, _ := json.Marshal(total)
	utils.Set(cacheKey, string(b), 5*time.Minute)

	log.Printf("[CloserX] Saved cache | %s (ttl=5m)", cacheKey)
	log.Printf("[CloserX] Request finished in %s", time.Since(reqStart))

	utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved successfully", total)
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
