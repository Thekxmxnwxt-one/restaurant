package routes

import (
	"restaurant/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTableRoutes(rg *gin.RouterGroup, h *handlers.TableHandlers) {
	table := rg.Group("/table")
	{
		table.GET("/", h.GetAllTable)
		table.GET("/:id", h.GetTableById)
		table.POST("/", h.AddTable)
		table.PATCH("/:id", h.UpdateTable)
		table.DELETE("/:id", h.DeleteTable)
	}
}
