package handler

import (
	"net/http"
	"sk-go-be/internal/model"
	"sk-go-be/internal/service"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	BrandService service.BrandService
}

func NewBrandHandler(brandService service.BrandService) *BrandHandler {
	return &BrandHandler{BrandService: brandService}
}

func (h *BrandHandler) GetBrands(c *gin.Context) {
	brands, err := h.BrandService.GetBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func (h *BrandHandler) GetBrandByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	brand, err := h.BrandService.GetBrandByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func (h *BrandHandler) CreateNewBrand(c *gin.Context) {
	var brand model.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.BrandService.CreateNewBrand(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, brand)
}

func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	var brand model.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.BrandService.UpdateBrand(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.BrandService.DeleteBrand(uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil) // No content response for successful deletion
}
