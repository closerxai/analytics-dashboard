package stripeclient

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"backend/utils"

	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/balancetransaction"
	"github.com/stripe/stripe-go/v84/creditnote"
	"github.com/stripe/stripe-go/v84/dispute"
	"github.com/stripe/stripe-go/v84/invoice"
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

	log.Printf("[Stripe] Fetching BalanceTransactions for key %s | start=%s end=%s", c.SecretKey, startDate, endDate)

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
		log.Printf("[Stripe] ERROR reading balance transactions for key %s: %v", c.SecretKey, err)
		return 0, 0, 0, err
	}

	log.Printf("[Stripe] Totals for key %s: revenue=%d refunded=%d disputes=%d",
		c.SecretKey,
		revenue, refunded, disputes)

	return revenue, refunded, disputes, nil
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

	log.Printf("[Stripe] Monthly stats started for key %s | year=%d", c.SecretKey, year)

	var results []MonthlyStats

	for m := 1; m <= 12; m++ {

		start := firstOfMonth(year, time.Month(m))
		end := endOfMonth(start)

		startStr := start.Format("2006-01-02")
		endStr := end.Format("2006-01-02")

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

type Bucket struct {
	Count    int64
	Revenue  int64 // cents
	Refunded int64
	Disputes int64
	DisputeCount int64
	Profit   int64
}

type IAResult struct {
	Subscriptions map[string]*Bucket
	Credits       map[string]*Bucket
	Others        map[string]*Bucket
}

func initBucket(m map[string]*Bucket, key string) {
	if _, ok := m[key]; !ok {
		m[key] = &Bucket{}
	}
}

func (c *Client) GetInvoiceBasedIA(startDate, endDate string) (*IAResult, error) {

	stripe.Key = c.SecretKey
	log.Printf("[Stripe] GetInvoiceBasedIA started for key %s | start=%s end=%s", c.SecretKey, startDate, endDate)

	result := &IAResult{
		Subscriptions: map[string]*Bucket{},
		Credits:       map[string]*Bucket{},
		Others:        map[string]*Bucket{},
	}

	// -------------------------------------------------
	// 1️⃣ INVOICES = REVENUE
	// -------------------------------------------------
	ip := &stripe.InvoiceListParams{
		Status: stripe.String("paid"),
	}
	ip.Limit = stripe.Int64(1000)
	ip.AddExpand("data.lines")

	if startDate != "" {
		ip.CreatedRange = &stripe.RangeQueryParams{
			GreaterThanOrEqual: parseTimestamp(startDate),
		}
	}
	if endDate != "" {
		if ip.CreatedRange == nil {
			ip.CreatedRange = &stripe.RangeQueryParams{}
		}
		ip.CreatedRange.LesserThanOrEqual = parseTimestamp(endDate)
	}

	invIter := invoice.List(ip)

	for invIter.Next() {
		inv := invIter.Invoice()

		for _, line := range inv.Lines.Data {

			amount := line.Amount // cents
			key := fmt.Sprintf("%.2f", float64(amount)/100)

			desc := strings.ToLower(line.Description)

			isSubscription := inv.BillingReason == stripe.InvoiceBillingReasonSubscriptionCycle &&
				line.Pricing != nil &&
				line.Pricing.PriceDetails != nil

			// ---- SUBSCRIPTIONS ----
			if isSubscription {
				initBucket(result.Subscriptions, key)
				b := result.Subscriptions[key]
				b.Count++
				b.Revenue += amount
				continue
			}

			// ---- CREDITS ----
			if strings.Contains(desc, "credit") {
				initBucket(result.Credits, key)
				b := result.Credits[key]
				b.Count++
				b.Revenue += amount
				continue
			}

			// ---- OTHERS ----
			initBucket(result.Others, key)
			b := result.Others[key]
			b.Count++
			b.Revenue += amount
		}
	}

	if err := invIter.Err(); err != nil {
		return nil, err
	}

	// -------------------------------------------------
	// 2️⃣ CREDIT NOTES = REFUNDS
	// -------------------------------------------------
	cnp := &stripe.CreditNoteListParams{}
	cnp.Limit = stripe.Int64(1000)

	if startDate != "" {
		cnp.CreatedRange = &stripe.RangeQueryParams{
			GreaterThanOrEqual: parseTimestamp(startDate),
		}
	}

	cnIter := creditnote.List(cnp)

	for cnIter.Next() {
		cn := cnIter.CreditNote()

		for _, line := range cn.Lines.Data {
			amount := utils.Abs64(line.Amount)
			key := fmt.Sprintf("%.2f", float64(amount)/100)

			for _, section := range []map[string]*Bucket{
				result.Subscriptions,
				result.Credits,
				result.Others,
			} {
				if b, ok := section[key]; ok {
					b.Refunded += amount
					break
				}
			}
		}
	}

	if err := cnIter.Err(); err != nil {
		return nil, err
	}

	// -------------------------------------------------
	// 3️⃣ DISPUTES (only lost disputes)
	// -------------------------------------------------
	dp := &stripe.DisputeListParams{}
	dp.Limit = stripe.Int64(1000)

	if startDate != "" {
		dp.Filters.AddFilter("created[gte]", "",
			fmt.Sprintf("%d", parseTimestamp(startDate)))
	}

	if endDate != "" {
		dp.Filters.AddFilter("created[lte]", "",
			fmt.Sprintf("%d", parseTimestamp(endDate)))
	}

	log.Print("[Stripe] Disputes parameters:", dp)
	dIter := dispute.List(dp)
	log.Print("[Stripe] Disputes found:", dIter)

	for dIter.Next() {
		d := dIter.Dispute()
		fmt.Println("--------------------------------")
		fmt.Println(d.ID, d.Status, d.Amount)

		// only include disputes that were lost
		if d.Status != stripe.DisputeStatusLost {
			fmt.Println("Dispute not lost:", d.ID, d.Status, d.Amount)
			continue
		}

		amount := d.Amount
		fmt.Println("Dispute amount:", amount)
		key := fmt.Sprintf("%.2f", float64(amount)/100)
		fmt.Println("Dispute key:", key)

		for _, section := range []map[string]*Bucket{
			result.Subscriptions,
			result.Credits,
			result.Others,
		} {
			fmt.Println("Section:", section)
			if b, ok := section[key]; ok {
				fmt.Println("Bucket:", b)
				fmt.Println("Adding to bucket:", amount)
				fmt.Println("Existing disputes:", b.Disputes)
				fmt.Println("Existing dispute count:", b.DisputeCount)
				b.Disputes += amount
				b.DisputeCount++
				fmt.Println("New disputes:", b.Disputes)
				fmt.Println("New dispute count:", b.DisputeCount)
				break
			}
		}
		fmt.Println("--------------------------------")
	}

	if err := dIter.Err(); err != nil {
		return nil, err
	}

	// -------------------------------------------------
	// 4️⃣ PROFIT
	// -------------------------------------------------
	for _, section := range []map[string]*Bucket{
		result.Subscriptions,
		result.Credits,
		result.Others,
	} {
		for _, b := range section {
			b.Profit = b.Revenue - b.Refunded - b.Disputes
		}
	}

	return result, nil
}

type Metric struct {
	Count int64 `json:"count"`
	Total int64 `json:"total"`
}

type Totals struct {
	Revenue  Metric `json:"revenue"`
	Refunded Metric `json:"refunded"`
	Disputed Metric `json:"disputed"`
	Profit   Metric `json:"profit"`
}

type PriceBreakdown struct {
	Revenue  Metric `json:"revenue"`
	Refunded Metric `json:"refunded"`
	Disputed Metric `json:"disputed"`
	Profit   Metric `json:"profit"`
}

type Section struct {
	Total  Totals                    `json:"total"`
	Prices map[string]PriceBreakdown `json:"prices"`
}

type IAResponse struct {
	Total         Totals  `json:"total"`
	Subscriptions Section `json:"subscriptions"`
	Credits       Section `json:"credits"`
	Others        Section `json:"others"`
}

func FormatIAResponse(ia *IAResult) IAResponse {

	resp := IAResponse{
		Subscriptions: Section{Prices: map[string]PriceBreakdown{}},
		Credits:       Section{Prices: map[string]PriceBreakdown{}},
		Others:        Section{Prices: map[string]PriceBreakdown{}},
	}

	apply := func(src map[string]*Bucket, section *Section) {
		for price, b := range src {

			// profit count == revenue count (derived metric)
			profitCount := b.Count

			section.Prices[price] = PriceBreakdown{
				Revenue: Metric{
					Count: b.Count,
					Total: b.Revenue,
				},
				Refunded: Metric{
					Count: utils.BoolToInt64(b.Refunded > 0),
					Total: b.Refunded,
				},
				Disputed: Metric{
					Count: b.DisputeCount,
					Total: b.Disputes,
				},
				Profit: Metric{
					Count: profitCount,
					Total: b.Profit,
				},
			}

			// ---- section totals ----
			section.Total.Revenue.Count += b.Count
			section.Total.Revenue.Total += b.Revenue

			section.Total.Refunded.Count += utils.BoolToInt64(b.Refunded > 0)
			section.Total.Refunded.Total += b.Refunded

			section.Total.Disputed.Count += utils.BoolToInt64(b.Disputes > 0)
			section.Total.Disputed.Total += b.Disputes

			section.Total.Profit.Count += profitCount
			section.Total.Profit.Total += b.Profit

			// ---- global totals ----
			resp.Total.Revenue.Count += b.Count
			resp.Total.Revenue.Total += b.Revenue

			resp.Total.Refunded.Count += utils.BoolToInt64(b.Refunded > 0)
			resp.Total.Refunded.Total += b.Refunded

			resp.Total.Disputed.Count += utils.BoolToInt64(b.Disputes > 0)
			resp.Total.Disputed.Total += b.Disputes

			resp.Total.Profit.Count += profitCount
			resp.Total.Profit.Total += b.Profit
		}
	}

	apply(ia.Subscriptions, &resp.Subscriptions)
	apply(ia.Credits, &resp.Credits)
	apply(ia.Others, &resp.Others)

	return resp
}
