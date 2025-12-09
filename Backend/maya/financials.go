package maya

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
	Revenue      int64 `json:"revenue"`
	Refunded     int64 `json:"refunded"`
	DisputesLost int64 `json:"disputes_lost"`
	Profit       int64 `json:"profit"`
}

// Global client instance for Maya
var client *stripeclient.Client

func Init() {
	key := os.Getenv("MAYA_STRIPE_KEY")
	client = stripeclient.New(key)
}

// ⚡ Fetch stats in parallel using your new Client method
func fetchStatsParallel(client *stripeclient.Client, startDate, endDate string) (int64, int64, int64, error) {
	return client.FetchStatsInParallel(startDate, endDate)
}

func GetFinancialStats(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate, endDate = utils.ApplyDefaultMonth(startDate, endDate)
	}

	log.Printf("[Maya] Request received | start_date=%s end_date=%s", startDate, endDate)

	if client == nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Stripe client not initialized", nil)
		return
	}

	// Cache key
	cacheKey := "maya:" + startDate + ":" + endDate

	// 1️⃣ Try cache
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data FinancialStats
		json.Unmarshal([]byte(cached), &data)

		utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved (cache)", data)
		return
	}

	// 2️⃣ Fetch in parallel using Stripe goroutines
	revenue, refunded, disputesLost, err := fetchStatsParallel(client, startDate, endDate)
	if err != nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch financial data", nil)
		return
	}

	total := FinancialStats{
		Revenue:      revenue,
		Refunded:     refunded,
		DisputesLost: disputesLost,
		Profit:       revenue - refunded - disputesLost,
	}

	// 3️⃣ Cache for 5 minutes
	b, _ := json.Marshal(total)
	utils.Set(cacheKey, string(b), 5*time.Minute)

	// 4️⃣ Return response
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

	if client == nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Stripe client not initialized", nil)
		return
	}

	cacheKey := "maya_monthly:" + yearStr

	// 1️⃣ Try cache
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data []stripeclient.MonthlyStats
		json.Unmarshal([]byte(cached), &data)

		utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved (cache)", data)
		return
	}

	// 2️⃣ Fetch monthly stats (parallel in Stripe client)
	monthly, err := client.GetMonthlyStats(year)
	if err != nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch monthly stats", nil)
		return
	}

	// 3️⃣ Cache for 30 minutes
	b, _ := json.Marshal(monthly)
	utils.Set(cacheKey, string(b), 30*time.Minute)

	// 4️⃣ Return final data
	utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved successfully", monthly)
}
