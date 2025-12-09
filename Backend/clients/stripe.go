package stripeclient

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/balancetransaction"
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

func parseTimestamp(date string) int64 {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("[Stripe] WARN invalid date %s", date)
		return 0
	}
	return t.Unix()
}

// --------------------------------------------------------------
// FAST: Fetch Revenue, Refunds, Disputes via Balance Transactions
// --------------------------------------------------------------
func (c *Client) GetTotals(startDate, endDate string) (revenue int64, refunded int64, disputes int64, err error) {

	stripe.Key = c.SecretKey

	log.Printf("[Stripe] Fetching BalanceTransactions | start=%s end=%s", startDate, endDate)

	params := &stripe.BalanceTransactionListParams{}
	params.Limit = stripe.Int64(1000) // large pages => fewer requests
	params.CreatedRange = &stripe.RangeQueryParams{}

	if startDate != "" {
		params.CreatedRange.GreaterThanOrEqual = parseTimestamp(startDate)
	}
	if endDate != "" {
		params.CreatedRange.LesserThanOrEqual = parseTimestamp(endDate)
	}

	iter := balancetransaction.List(params)

	for iter.Next() {
		bt := iter.BalanceTransaction()

		switch bt.Type {

		case "charge":
			// Money received (positive)
			revenue += bt.Amount

		case "refund":
			// Money sent back (negative balance, but Stripe returns positive amount)
			refunded += bt.Amount

		case "dispute":
			// Dispute amounts are already positive = lost amount
			disputes += bt.Amount
		}
	}

	if err := iter.Err(); err != nil {
		log.Printf("[Stripe] ERROR reading balance transactions: %v", err)
		return 0, 0, 0, err
	}

	log.Printf("[Stripe] Totals: revenue=%d refunded=%d disputes=%d",
		revenue, refunded, disputes)

	return revenue, refunded, disputes, nil
}

// --------------------------------------------------------------
// PARALLEL WRAPPER (same API as before)
// --------------------------------------------------------------
func (c *Client) FetchStatsInParallel(startDate, endDate string) (int64, int64, int64, error) {

	log.Printf("[Stripe] Parallel fetch started | start=%s end=%s", startDate, endDate)

	var wg sync.WaitGroup
	wg.Add(1)

	var revenue, refunded, disputes int64
	var err error

	go func() {
		defer wg.Done()
		revenue, refunded, disputes, err = c.GetTotals(startDate, endDate)
	}()

	wg.Wait()
	return revenue, refunded, disputes, err
}

// --------------------------------------------------------------
// MONTHLY STATS (now super fast)
// --------------------------------------------------------------
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

	log.Printf("[Stripe] Monthly stats started | year=%d", year)

	var results []MonthlyStats

	for m := 1; m <= 12; m++ {

		start := firstOfMonth(year, time.Month(m))
		end := endOfMonth(start)

		startStr := start.Format("2006-01-02")
		endStr := end.Format("2006-01-02")

		log.Printf("[Stripe] Fetch month=%s | %s → %s", start.Format("2006-01"), startStr, endStr)

		revenue, refunded, disputes, err := c.GetTotals(startStr, endStr)
		if err != nil {
			return nil, err
		}

		results = append(results, MonthlyStats{
			Month:   start.Format("2006-01"),
			Revenue: revenue,
			Profit:  revenue - refunded - disputes,
		})
	}

	return results, nil
}
