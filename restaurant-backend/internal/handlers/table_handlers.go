package handlers

import (
	"net/http"
	"restaurant/internal/models"
	"restaurant/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TableHandlers struct {
	rs *service.TableApp
}

func NewTableHandlers(rs *service.TableApp) *TableHandlers {
	return &TableHandlers{rs: rs}
}

func (h *TableHandlers) GetAllTable(c *gin.Context) {
	menus, err := h.rs.GetAllTable(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (h *TableHandlers) GetTableById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	table, err := h.rs.GetTableById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)
}

func (h *TableHandlers) AddTable(c *gin.Context) {
	var table models.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.rs.AddTable(c.Request.Context(), &table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, table)
}

func (h *TableHandlers) UpdateTable(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	type UpdateTableInput struct {
		Name     *string `json:"name"`
		Capacity *int    `json:"capacity"`
		Status   *string `json:"status"`
	}

	var input UpdateTableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = h.rs.UpdateTable(
		c.Request.Context(),
		id,
		input.Name,
		input.Capacity,
		input.Status,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "table updated successfully"})
}

func (h *TableHandlers) DeleteTable(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.rs.DeleteTable(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "table deleted successfully"})
}
