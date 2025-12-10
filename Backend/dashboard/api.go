package dashboard

import "github.com/gin-gonic/gin"

func HealthAPI(c *gin.Context) {
	HealthCheckAll(c)
}
