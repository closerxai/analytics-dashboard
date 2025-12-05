package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/internal/handlers"
	"backend/internal/services"
	stripeclient "backend/internal/stripe"
)

func main() {
	godotenv.Load()
	stripeclient.Init()

	client := stripeclient.New()
	srv := services.NewAnalyticsService(client)
	handler := handlers.NewAnalyticsHandler(srv)

	r := gin.Default()

	r.GET("/analytics", handler.GetAnalytics)

	log.Println("Server running at :8080")
	r.Run(":8080")
}
