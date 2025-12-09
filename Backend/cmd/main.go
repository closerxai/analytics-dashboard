package main

import (
	"backend/closerx"
	"backend/maya"
	"backend/snowie"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	closerxGroup := router.Group("/api/closerx")
	closerxGroup.GET("/financials", closerx.GetFinancialStats)
	closerxGroup.GET("/financials/graph", closerx.GetMonthlyStats)

	mayaGroup := router.Group("/api/maya")
	mayaGroup.GET("/financials", maya.GetFinancialStats)
	mayaGroup.GET("/financials/graph", maya.GetMonthlyStats)

	snowieGroup := router.Group("/api/snowie")
	snowieGroup.GET("/financials", snowie.GetFinancialStats)
	snowieGroup.GET("/financials/graph", snowie.GetMonthlyStats)

	router.Run(":8080")
}
