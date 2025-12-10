package infra

import "github.com/gin-gonic/gin"

func AWSBillingAPI(c *gin.Context) {
	GetAWSBilling(c)
}
