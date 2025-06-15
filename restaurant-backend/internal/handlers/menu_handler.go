package handlers

import (
	"net/http"
	"restaurant/internal/models"
	"restaurant/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuHandlers struct {
	rs *service.MenuApp
}

func NewMenuHandlers(rs *service.MenuApp) *MenuHandlers {
	return &MenuHandlers{rs: rs}
}

func (h *MenuHandlers) GetAllMenu(c *gin.Context) {
	menus, err := h.rs.GetAllMenu(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (h *MenuHandlers) GetMenuById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	menu, err := h.rs.GetMenuById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func (h *MenuHandlers) AddMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.rs.AddMenu(c.Request.Context(), &menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}

func (h *MenuHandlers) UpdateMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	type UpdateMenuInput struct {
		Name        *string  `json:"name"`
		Image       *string  `json:"image_url"`
		Description *string  `json:"description"`
		Price       *float64 `json:"price"`
		Category    *string  `json:"category"`
		Available   *bool    `json:"available"`
	}

	var input UpdateMenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = h.rs.UpdateMenu(
		c.Request.Context(),
		id,
		input.Name,
		input.Image,
		input.Description,
		input.Price,
		input.Category,
		input.Available,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu updated successfully"})
}

func (h *MenuHandlers) DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.rs.DeleteMenu(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "menu deleted successfully"})
}
