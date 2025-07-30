package handler

import (
	"net/http"
	"sk-go-be/internal/model"
	"sk-go-be/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

func (h *CategoryHandler) GetProductCategoryList(c *gin.Context) {
	categories, err := h.CategoryService.GetProductCategoryList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func (h *CategoryHandler) CreateNewProductCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.CategoryService.CreateNewProductCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) UpdateProductCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paramUuid := c.Param("uuid") // Assuming UUID is passed as a URL parameter
	formattedUUID, err := uuid.Parse(paramUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.UUID = formattedUUID

	if err := h.CategoryService.UpdateProductCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// Optional: If you want to support deletion of categories
func (h *CategoryHandler) DeleteProductCategory(c *gin.Context) {
	paramUuid := c.Param("uuid") // Assuming UUID is passed as a URL parameter
	if err := h.CategoryService.DeleteProductCategory(paramUuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil) // No content response for successful deletion
}
