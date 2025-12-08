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
