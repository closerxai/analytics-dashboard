package dashboard

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"backend/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	ceTypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"

	"github.com/gin-gonic/gin"
)

func GetAWSBilling(c *gin.Context) {
	ctx := context.TODO()

	// Try loading AWS credentials
	cfg, err := config.LoadDefaultConfig(ctx)
	fmt.Println(cfg)
	

	// Real AWS call
	ce := costexplorer.NewFromConfig(cfg)

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := now

	res, err := ce.GetCostAndUsage(ctx, &costexplorer.GetCostAndUsageInput{
		Metrics: []string{"UnblendedCost"},
		TimePeriod: &ceTypes.DateInterval{
			Start: aws.String(start.Format("2006-01-02")),
			End:   aws.String(end.Format("2006-01-02")),
		},

		Granularity: "MONTHLY",
	})

	// If AWS query fails
	if err != nil {
		utils.CustomResponse(c, http.StatusInternalServerError, false, "Failed to fetch AWS billing", gin.H{
			"error": err.Error(),
		})
		return
	}

	// Extract cost
	total := "0"
	if len(res.ResultsByTime) > 0 {
		amount := res.ResultsByTime[0].Total["UnblendedCost"].Amount
		if amount != nil {
			total = *amount
		}
	}

	// Success response
	utils.CustomResponse(c, http.StatusOK, true, "AWS billing fetched", gin.H{
		"totalCost": total,
		"start":     start.Format("2006-01-02"),
		"end":       end.Format("2006-01-02"),
	})
}
