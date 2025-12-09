package stripeclient

import (
	"os"
	"time"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/charge"
	"github.com/stripe/stripe-go/v80/dispute"
	"github.com/stripe/stripe-go/v80/refund"
	"sync"
)

func Init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

type Client struct {
	SecretKey string
}

func New(secret string) *Client {
	return &Client{SecretKey: secret}
}

func (c *Client) GetRevenue(startDate, endDate string) (int64, error) {
	stripe.Key = c.SecretKey
	params := &stripe.ChargeListParams{}
	params.Limit = stripe.Int64(100)
	params.CreatedRange = &stripe.RangeQueryParams{}
	if startDate != "" {
		params.CreatedRange.GreaterThanOrEqual = parseTimestamp(startDate)
	}
	if endDate != "" {
		params.CreatedRange.LesserThanOrEqual = parseTimestamp(endDate)
	}

	iter := charge.List(params)

	var total int64 = 0

	for iter.Next() {
		ch := iter.Charge()
		if ch.Paid && ch.Status == "succeeded" {
			total += ch.Amount
		}
	}

	return total, iter.Err()
}

func (c *Client) GetRefunded(startDate, endDate string) (int64, error) {
	stripe.Key = c.SecretKey
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(100)

	params.CreatedRange = &stripe.RangeQueryParams{}

	if startDate != "" {
		params.CreatedRange.GreaterThanOrEqual = parseTimestamp(startDate)
	}
	if endDate != "" {
		params.CreatedRange.LesserThanOrEqual = parseTimestamp(endDate)
	}

	iter := refund.List(params)

	var total int64 = 0
	for iter.Next() {
		total += iter.Refund().Amount
	}

	return total, iter.Err()
}

func (c *Client) GetDisputesLost(startDate, endDate string) (int64, error) {
	stripe.Key = c.SecretKey

	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(100)

	// Proper date range filtering using CreatedRange
	params.CreatedRange = &stripe.RangeQueryParams{}

	if startDate != "" {
		params.CreatedRange.GreaterThanOrEqual = parseTimestamp(startDate)
	}
	if endDate != "" {
		params.CreatedRange.LesserThanOrEqual = parseTimestamp(endDate)
	}

	iter := dispute.List(params)

	var total int64
	for iter.Next() {
		d := iter.Dispute()
		if d.Status == stripe.DisputeStatusLost {
			total += d.Amount
		}
	}

	return total, iter.Err()
}

// Helper function to convert date string to Unix timestamp
func parseTimestamp(date string) int64 {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func (c *Client) FetchStatsInParallel(startDate, endDate string) (int64, int64, int64, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	var revenue, refunded, disputes int64
	var err1, err2, err3 error

	go func() {
		defer wg.Done()
		revenue, err1 = c.GetRevenue(startDate, endDate)
	}()

	go func() {
		defer wg.Done()
		refunded, err2 = c.GetRefunded(startDate, endDate)
	}()

	go func() {
		defer wg.Done()
		disputes, err3 = c.GetDisputesLost(startDate, endDate)
	}()

	wg.Wait()

	if err1 != nil {
		return 0, 0, 0, err1
	}
	if err2 != nil {
		return 0, 0, 0, err2
	}
	if err3 != nil {
		return 0, 0, 0, err3
	}

	return revenue, refunded, disputes, nil
}


type MonthlyStats struct {
	Month        string `json:"month"`  // "2024-01"
	Revenue      int64  `json:"revenue"`
	Profit       int64  `json:"profit"`
}


// Get the first day of the month (YYYY-MM)
func firstOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// Get the last day of month
func endOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 0).Add(-time.Second) // last second of month
}


func (c *Client) GetMonthlyStats(year int) ([]MonthlyStats, error) {
	var results []MonthlyStats

	for month := 1; month <= 12; month++ {
		start := firstOfMonth(year, time.Month(month))
		end := endOfMonth(start)

		startStr := start.Format("2006-01-02")
		endStr := end.Format("2006-01-02")

		// Use your fast parallel fetch
		revenue, refunded, disputes, err := c.FetchStatsInParallel(startStr, endStr)
		if err != nil {
			return nil, err
		}

		stats := MonthlyStats{
			Month:        start.Format("2006-01"), // graph-friendly
			Revenue:      revenue,
			Profit:       revenue - refunded - disputes,
		}

		results = append(results, stats)
	}

	return results, nil
}
