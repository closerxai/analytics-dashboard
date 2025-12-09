package main

import (
	"backend/closerx"
	"backend/maya"
	"backend/snowie"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))


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
