package handler

import (
	"sk-go-be/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartService service.CartService
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{CartService: cartService}
}

func (h *CartHandler) GetShoppingCartByUUID(c *gin.Context) {
	// Get the UUID from the URL
	uuid := c.Param("uuid")

	// Get the cart from the service
	cart, err := h.CartService.GetShoppingCartByUUID(uuid)
	if err != nil {
		// Handle error
		c.JSON(404, gin.H{"error": "Cart not found"})
		return
	}

	// Return the cart as JSON
	c.JSON(200, cart)
}

func (h *CartHandler) GetShoppingCartByUserUUID(c *gin.Context) {
	// Get the UUID from the URL
	uuid := c.Param("uuid")

	// Get the cart from the service
	cart, err := h.CartService.GetShoppingCartByUserUUID(uuid)
	if err != nil {
		// Handle error
		c.JSON(404, gin.H{"error": "Cart not found"})
		return
	}

	// Return the cart as JSON
	c.JSON(200, cart)
}
