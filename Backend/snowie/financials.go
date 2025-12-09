package snowie

import (
	stripeclient "backend/clients"
	"backend/utils"
	"net/http"
	"os"
	"log"

	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FinancialStats struct {
	Revenue      int64 `json:"revenue"`
	Refunded     int64 `json:"refunded"`
	DisputesLost int64 `json:"disputes_lost"`
	Profit       int64 `json:"profit"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

func fetchForAccount(key, startDate, endDate string, ch chan<- FinancialStats, errCh chan<- error) {
	client := stripeclient.New(key)

	r, f, d, err := client.FetchStatsInParallel(startDate, endDate)
	if err != nil {
		errCh <- err
		return
	}

	ch <- FinancialStats{
		Revenue:      r,
		Refunded:     f,
		DisputesLost: d,
		Profit:       r - f - d,
		StartDate:    startDate,
		EndDate:      endDate,
	}
}

func GetFinancialStats(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate, endDate = utils.ApplyDefaultMonth(startDate, endDate)
	}

	log.Printf("[Snowie] Request received | start_date=%s end_date=%s", startDate, endDate)

	cacheKey := "snowie:" + startDate + ":" + endDate

	// 1. Try cache
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data FinancialStats
		json.Unmarshal([]byte(cached), &data)

		utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved (cache)", data)
		return
	}

	// 2. Parallel fetch from Stripe keys
	keys := []string{
		os.Getenv("SNOWIE_STRIPE_KEY_1"),
		os.Getenv("SNOWIE_STRIPE_KEY_2"),
	}

	statsCh := make(chan FinancialStats, len(keys))
	errCh := make(chan error, len(keys))

	for _, key := range keys {
		go fetchForAccount(key, startDate, endDate, statsCh, errCh)
	}

	var total FinancialStats

	// 3. Wait for results
	for range keys {
		select {
		case s := <-statsCh:
			total.Revenue += s.Revenue
			total.Refunded += s.Refunded
			total.DisputesLost += s.DisputesLost
		case err := <-errCh:
			utils.CustomResponse(c, http.StatusInternalServerError, false, err.Error(), nil)
			return
		}
	}

	total.Profit = total.Revenue - total.Refunded - total.DisputesLost
	total.StartDate = startDate
	total.EndDate = endDate

	// 4. Write to cache (5 minutes)
	b, _ := json.Marshal(total)
	utils.Set(cacheKey, string(b), 5*time.Minute)

	// 5. Return response
	utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved successfully", total)
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
