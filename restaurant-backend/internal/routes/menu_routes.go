package routes

import (
	"restaurant/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(rg *gin.RouterGroup, h *handlers.MenuHandlers) {
	menu := rg.Group("/menu")
	{
		menu.GET("/", h.GetAllMenu)
		menu.GET("/:id", h.GetMenuById)
		menu.POST("/", h.AddMenu)
		menu.PATCH("/:id", h.UpdateMenu)
		menu.DELETE("/:id", h.DeleteMenu)
	}
}
