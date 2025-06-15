package handlers

import (
	"net/http"
	"restaurant/internal/service"

	"github.com/gin-gonic/gin"
)

type RestaurantHandlers struct {
	rs *service.RestaurantApp
}

func NewRestaurantHandlers(rs *service.RestaurantApp) *RestaurantHandlers {
	return &RestaurantHandlers{rs: rs}
}

func (h *RestaurantHandlers) HealthCheck(c *gin.Context) {
	err := h.rs.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"reason": "Database connection failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
