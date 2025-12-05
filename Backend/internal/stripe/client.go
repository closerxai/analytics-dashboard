package stripeclient

import (
	"os"

	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/charge"
	"github.com/stripe/stripe-go/v80/refund"
	"github.com/stripe/stripe-go/v80/dispute"
)

func Init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) GetRevenue() (int64, error) {
	params := &stripe.ChargeListParams{}
	params.Limit = stripe.Int64(100)

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


func (c *Client) GetRefunded() (int64, error) {
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(100) // ✔ correct

	iter := refund.List(params)

	var total int64 = 0
	for iter.Next() {
		total += iter.Refund().Amount
	}

	return total, iter.Err()
}


func (c *Client) GetDisputesLost() (int64, error) {
	params := &stripe.DisputeListParams{}
	params.Limit = stripe.Int64(100) // ✔ correct

	iter := dispute.List(params)

	var total int64 = 0
	for iter.Next() {
		d := iter.Dispute()
		if d.Status == stripe.DisputeStatusLost {
			total += d.Amount
		}
	}

	return total, iter.Err()
}

