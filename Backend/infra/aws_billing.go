package infra

import (
	"context"
	"time"

	"backend/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	ceTypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// ------------ AWS CLIENT ------------
func NewAWSClient(ctx context.Context) (*costexplorer.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return costexplorer.NewFromConfig(cfg), nil
}

// ------------ RESTY CLIENT (you may use later for external APIs) ------------
func NewRestyClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	return client
}

// ------------ MAIN FUNCTION ------------
func GetAWSBilling(c *gin.Context) {
	ctx := context.TODO()

	// AWS client
	ceClient, err := NewAWSClient(ctx)
	if err != nil {
		utils.CustomResponse(c, 500, false, "Failed to configure AWS", gin.H{
			"error": err.Error(),
		})
		return
	}

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := now

	// AWS Cost Explorer request
	resp, err := ceClient.GetCostAndUsage(ctx, &costexplorer.GetCostAndUsageInput{
		Metrics: []string{"UnblendedCost"},
		TimePeriod: &ceTypes.DateInterval{
			Start: aws.String(start.Format("2006-01-02")),
			End:   aws.String(end.Format("2006-01-02")),
		},
		Granularity: "MONTHLY",
	})

	if err != nil {
		utils.CustomResponse(c, 500, false, "Failed to fetch AWS Billing", gin.H{
			"error": err.Error(),
		})
		return
	}

	// Extract total cost
	total := "0"
	if len(resp.ResultsByTime) > 0 {
		amount := resp.ResultsByTime[0].Total["UnblendedCost"].Amount
		if amount != nil {
			total = *amount
		}
	}

	utils.CustomResponse(c, 200, true, "AWS billing fetched", gin.H{
		"totalCost": total,
		"start":     start.Format("2006-01-02"),
		"end":       end.Format("2006-01-02"),
	})
}
