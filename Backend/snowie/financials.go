package snowie

import (
	stripeclient "backend/clients"
	"backend/utils"
	"net/http"
	"os"

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
}

// Global client instance for Maya
var client *stripeclient.Client

func Init() {
	key := os.Getenv("SNOWIE_STRIPE_KEY")
	client = stripeclient.New(key)
}

// ⚡ Fetch stats in parallel using your new Client method
func fetchStatsParallel(client *stripeclient.Client, startDate, endDate string) (int64, int64, int64, error) {
	return client.FetchStatsInParallel(startDate, endDate)
}

func GetFinancialStats(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	cacheKey := "snowie:" + startDate + ":" + endDate

	// 1️⃣ Try Redis cache
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data FinancialStats
		json.Unmarshal([]byte(cached), &data)

		utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved (cache)", data)
		return
	}

	// 3️⃣ Fetch in parallel using your FetchStatsInParallel method
	revenue, refunded, disputes, err := fetchStatsParallel(client, startDate, endDate)
	if err != nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch financial data", nil)
		return
	}

	total := FinancialStats{
		Revenue:      revenue,
		Refunded:     refunded,
		DisputesLost: disputes,
		Profit:       revenue - refunded - disputes,
	}

	// 4️⃣ Cache result for 5 minutes
	b, _ := json.Marshal(total)
	utils.Set(cacheKey, string(b), 5*time.Minute)

	// 5️⃣ Return response
	utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved successfully", total)
}

func GetMonthlyStats(c *gin.Context) {
	yearStr := c.Query("year")
	if yearStr == "" {
		yearStr = "2025"
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		utils.CustomResponse(c, http.StatusBadRequest, false, "Invalid year", nil)
		return
	}

	cacheKey := "snowie_monthly:" + yearStr

	// 1️⃣ Try cache
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data []stripeclient.MonthlyStats
		json.Unmarshal([]byte(cached), &data)
		utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved (cache)", data)
		return
	}

	// 2️⃣ Ensure Stripe client exists
	if client == nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Stripe client not initialized", nil)
		return
	}

	// 3️⃣ Fetch monthly stats (Stripe goroutine parallelized internally)
	monthly, err := client.GetMonthlyStats(year)
	if err != nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch monthly stats", nil)
		return
	}

	// 4️⃣ Cache output for 30 minutes (recommended for graphs)
	b, _ := json.Marshal(monthly)
	utils.Set(cacheKey, string(b), 30*time.Minute)

	// 5️⃣ Return results
	utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved successfully", monthly)
}
