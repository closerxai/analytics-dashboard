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
	data, err := h.srv.GetAnalytics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch analytics"})
		return
	}
	c.JSON(http.StatusOK, data)
}
