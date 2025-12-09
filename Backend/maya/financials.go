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
	Revenue      int64  `json:"revenue"`
	Refunded     int64  `json:"refunded"`
	DisputesLost int64  `json:"disputes_lost"`
	Profit       int64  `json:"profit"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

var client *stripeclient.Client

func Init() {
	secret := os.Getenv("MAYA_STRIPE_KEY")
	client = stripeclient.New(secret)
}

// ------------------------------
//       SINGLE ACCOUNT API
// ------------------------------

func GetFinancialStats(c *gin.Context) {
	start := time.Now()

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate, endDate = utils.ApplyDefaultMonth(startDate, endDate)
	}

	log.Printf("[Maya] Request | start=%s end=%s", startDate, endDate)

	if client == nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Stripe client not initialized", nil)
		return
	}

	cacheKey := "maya:" + startDate + ":" + endDate

	// 1️⃣ Cache check
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		log.Printf("[Maya] Cache HIT | %s", cacheKey)

		var data FinancialStats
		json.Unmarshal([]byte(cached), &data)
		utils.CustomResponse(c, http.StatusOK, true, "Financial stats (cache)", data)
		return
	}

	log.Printf("[Maya] Cache MISS | %s", cacheKey)

	// 2️⃣ FAST Stripe balance transaction fetch
	revenue, refunded, disputes, err := client.GetTotals(startDate, endDate)
	if err != nil {
		log.Printf("[Maya] ERROR fetch: %v", err)
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch financial data", nil)
		return
	}
	

	total := FinancialStats{
		Revenue:      revenue,
		Refunded:     refunded,
		DisputesLost: disputes,
		Profit:       revenue - refunded - disputes,
		StartDate:    startDate,
		EndDate:      endDate,
	}

	log.Printf("[Maya] Totals | revenue=%d refunded=%d disputes=%d profit=%d",
		total.Revenue, total.Refunded, total.DisputesLost, total.Profit)

	// 3️⃣ Cache
	bytes, _ := json.Marshal(total)
	utils.Set(cacheKey, string(bytes), 5*time.Minute)

	log.Printf("[Maya] Cached | %s", cacheKey)
	log.Printf("[Maya] Completed in %s", time.Since(start))

	utils.CustomResponse(c, http.StatusOK, true, "Financial stats retrieved successfully", total)
}

// ------------------------------
//          MONTHLY STATS
// ------------------------------

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

	// 1️⃣ Cache check
	if cached, err := utils.Get(cacheKey); err == nil && cached != "" {
		var data []stripeclient.MonthlyStats
		json.Unmarshal([]byte(cached), &data)
		utils.CustomResponse(c, http.StatusOK, true, "Monthly stats (cache)", data)
		return
	}

	// 2️⃣ FAST monthly stats
	monthlyStats, err := client.GetMonthlyStats(year)
	if err != nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch monthly stats", nil)
		return
	}

	// 3️⃣ Cache
	bytes, _ := json.Marshal(monthlyStats)
	utils.Set(cacheKey, string(bytes), 30*time.Minute)

	// 4️⃣ Return response
	utils.CustomResponse(c, http.StatusOK, true, "Monthly stats retrieved successfully", monthlyStats)
}
