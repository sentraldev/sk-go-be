package handler

import (
	"net/http"
	"sk-go-be/internal/service"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	WishlistService service.WishlistService
}

func NewWishlistHandler(wishlistService service.WishlistService) *WishlistHandler {
	return &WishlistHandler{WishlistService: wishlistService}
}

func (h *WishlistHandler) GetWishlistByUserUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	wishlist, err := h.WishlistService.GetWishlistByUserUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wishlist)
}
