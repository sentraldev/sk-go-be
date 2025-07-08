package handler

import (
	"net/http"
	"sk-go-be/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) GetProductByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	product, err := h.ProductService.GetProductByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
