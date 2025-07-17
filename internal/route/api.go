package route

import (
	"sk-go-be/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(
	r *gin.Engine,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	wishlistHandler *handler.WishlistHandler,
	discountHandler *handler.DiscountHandler,
) {
	v1 := r.Group("/api/v1")

	// Public API routes
	public := v1.Group("/public")
	// Add public endpoints here
	{
		// public.GET("/products", handler.ListUsers)
		SetupCartPublicRoutes(public, *cartHandler)
	}

	// Admin API routes (should be protected by JWT middleware)
	admin := v1.Group("/admin")
	// Add admin endpoints here
	{
		SetupCartAdminRoutes(admin, *cartHandler)
	}

}

func SetupCartPublicRoutes(r *gin.RouterGroup, handler handler.CartHandler) {
	r.GET("/cart/:uuid", handler.GetShoppingCartByUUID)
	r.GET("/cart/user/:uuid", handler.GetShoppingCartByUserUUID)
}

func SetupCartAdminRoutes(r *gin.RouterGroup, handler handler.CartHandler) {
	// r.GET("/cart/:uuid", handler.GetShoppingCartByUUID)
	// r.GET("/cart/user/:uuid", handler.GetShoppingCartByUserUUID)
}
