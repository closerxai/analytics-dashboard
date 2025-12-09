package handlers

import (
	"net/http"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	srv *services.AnalyticsService
}

func NewAnalyticsHandler(s *services.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{s}
}

func (h *AnalyticsHandler) GetAnalytics(c *gin.Context) {
	product := c.Query("product")
	if product == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product is required"})
		return
	}

	data, err := h.srv.GetAnalytics(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *AnalyticsHandler) GetRevenue(c *gin.Context) {
	product := c.Query("product")
	value, err := h.srv.GetRevenue(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"revenue": value})
}

func (h *AnalyticsHandler) GetRefunded(c *gin.Context) {
	product := c.Query("product")
	value, err := h.srv.GetRefunded(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"refunded": value})
}

func (h *AnalyticsHandler) GetDisputes(c *gin.Context) {
	product := c.Query("product")
	value, err := h.srv.GetDisputes(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"disputes": value})
}

func (h *AnalyticsHandler) GetProfit(c *gin.Context) {
	product := c.Query("product")
	value, err := h.srv.GetProfit(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"profit": value})
}
