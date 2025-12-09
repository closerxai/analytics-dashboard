package stripeclient

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/charge"
	"github.com/stripe/stripe-go/v80/dispute"
	"github.com/stripe/stripe-go/v80/refund"
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

	log.Printf("[Stripe] Fetching Revenue | start=%s end=%s", startDate, endDate)

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

	var total int64

	for iter.Next() {
		ch := iter.Charge()
		if ch.Paid && ch.Status == "succeeded" {
			total += ch.Amount
		}
	}

	log.Printf("[Stripe] Revenue Total = %d", total)
	return total, iter.Err()
}

func (c *Client) GetRefunded(startDate, endDate string) (int64, error) {
	stripe.Key = c.SecretKey

	log.Printf("[Stripe] Fetching Refunds | start=%s end=%s", startDate, endDate)

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

	var total int64

	for iter.Next() {
		r := iter.Refund()
		total += r.Amount
	}

	log.Printf("[Stripe] Refund Total = %d", total)
	return total, iter.Err()
}

func (c *Client) GetDisputesLost(startDate, endDate string) (int64, error) {
	stripe.Key = c.SecretKey

	log.Printf("[Stripe] Fetching DisputesLost | start=%s end=%s", startDate, endDate)

	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(100)
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

	log.Printf("[Stripe] DisputesLost Total = %d", total)
	return total, iter.Err()
}

func parseTimestamp(date string) int64 {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("[Stripe] WARN: invalid date '%s'", date)
		return 0
	}
	return t.Unix()
}

func (c *Client) FetchStatsInParallel(startDate, endDate string) (int64, int64, int64, error) {
	log.Printf("[Stripe] Parallel fetch started | start=%s end=%s", startDate, endDate)

	var wg sync.WaitGroup
	wg.Add(3)

	var revenue, refunded, disputes int64
	var err1, err2, err3 error

	go func() {
		defer wg.Done()
		revenue, err1 = c.GetRevenue(startDate, endDate)
		log.Printf("[Stripe] Revenue fetch done for key %s: %d", c.SecretKey, revenue)
	}()

	go func() {
		defer wg.Done()
		refunded, err2 = c.GetRefunded(startDate, endDate)
		log.Printf("[Stripe] Refund fetch done for key %s: %d", c.SecretKey, refunded)
	}()

	go func() {
		defer wg.Done()
		disputes, err3 = c.GetDisputesLost(startDate, endDate)
		log.Printf("[Stripe] Disputes fetch done for key %s: %d", c.SecretKey, disputes)
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

	log.Printf("[Stripe] Parallel totals for key %s | revenue=%d refunded=%d disputes=%d",
		c.SecretKey[len(c.SecretKey)-4:],
		revenue, refunded, disputes)

	return revenue, refunded, disputes, nil
}

type MonthlyStats struct {
	Month   string `json:"month"`
	Revenue int64  `json:"revenue"`
	Profit  int64  `json:"profit"`
}

func firstOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

func endOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 0).Add(-time.Second)
}

func (c *Client) GetMonthlyStats(year int) ([]MonthlyStats, error) {
	log.Printf("[Stripe] Monthly stats fetch started | year=%d", year)

	var results []MonthlyStats

	for m := 1; m <= 12; m++ {
		start := firstOfMonth(year, time.Month(m))
		end := endOfMonth(start)

		startStr := start.Format("2006-01-02")
		endStr := end.Format("2006-01-02")

		log.Printf("[Stripe] Fetching month=%s | start=%s end=%s",
			start.Format("2006-01"), startStr, endStr)

		revenue, refunded, disputes, err := c.FetchStatsInParallel(startStr, endStr)
		if err != nil {
			log.Printf("[Stripe] ERROR month=%s: %v", start.Format("2006-01"), err)
			return nil, err
		}

		stats := MonthlyStats{
			Month:   start.Format("2006-01"),
			Revenue: revenue,
			Profit:  revenue - refunded - disputes,
		}

		log.Printf("[Stripe] Month=%s totals | revenue=%d profit=%d",
			stats.Month, stats.Revenue, stats.Profit)

		results = append(results, stats)
	}

	return results, nil
}
