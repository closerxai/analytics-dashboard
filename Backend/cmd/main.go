package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/internal/handlers"
	"backend/internal/services"
	stripeclient "backend/clients"
)

func main() {
	godotenv.Load()

	stripeClients := map[string]*stripeclient.Client{
		"maya":    stripeclient.New(os.Getenv("MAYA_STRIPE_KEY")),
		"snowie":  stripeclient.New(os.Getenv("SNOWIE_STRIPE_KEY")),
		"closerx": stripeclient.New(os.Getenv("CLOSERX_STRIPE_KEY")),
	}

	srv := services.NewAnalyticsService(stripeClients)
	handler := handlers.NewAnalyticsHandler(srv)

	r := gin.Default()

	r.GET("/analytics", handler.GetAnalytics)
	r.GET("/analytics/revenue", handler.GetRevenue)
	r.GET("/analytics/refunded", handler.GetRefunded)
	r.GET("/analytics/disputes", handler.GetDisputes)
	r.GET("/analytics/profit", handler.GetProfit)

	log.Println("Server running at :8080")
	r.Run(":8080")
}
