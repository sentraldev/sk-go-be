package route

import (
	"sk-go-be/internal/handler"
	"sk-go-be/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(
	r *gin.Engine,
	fc *middleware.FirebaseClient,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	wishlistHandler *handler.WishlistHandler,
	discountHandler *handler.DiscountHandler,
	authHandler *handler.AuthHandler,
) {
	v1 := r.Group("/api/v1")

	// Public API routes
	public := v1.Group("/public")
	// Add public endpoints here
	{
		// public.GET("/products", handler.ListUsers)
		SetupProductRoutes(public, *productHandler)

		SetupCartPublicRoutes(public, *cartHandler)
		SetupAuthRoutes(public, *authHandler, *fc)
		SetupUserRoutes(public, *userHandler, *fc)
	}

	// Admin API routes (should be protected by JWT middleware)
	admin := v1.Group("/admin")
	admin.Use(fc.AuthMiddleware())
	// Add admin endpoints here
	{
		SetupCartAdminRoutes(admin, *cartHandler)
	}

}

func SetupCartPublicRoutes(r *gin.RouterGroup, handler handler.CartHandler) {
	r.GET("/cart/:uuid", handler.GetShoppingCartByUUID)
	r.GET("/cart/user/:uuid", handler.GetShoppingCartByUserUUID)
}

func SetupProductRoutes(r *gin.RouterGroup, handler handler.ProductHandler) {
	r.GET("/products/:uuid", handler.GetProductByUUID)
	r.GET("/products", handler.GetProducts)
}

func SetupCartAdminRoutes(r *gin.RouterGroup, handler handler.CartHandler) {
	// r.GET("/cart/:uuid", handler.GetShoppingCartByUUID)
	// r.GET("/cart/user/:uuid", handler.GetShoppingCartByUserUUID)
}

func SetupAuthRoutes(r *gin.RouterGroup, handler handler.AuthHandler, authMiddleware middleware.FirebaseClient) {
	r.Use(authMiddleware.AuthMiddleware()).POST("/register", handler.Register)
	// r.POST("/login", handler.Login) // Uncomment when login is implemented
}

func SetupUserRoutes(r *gin.RouterGroup, handler handler.UserHandler, authMiddleware middleware.FirebaseClient) {
	r.Use(authMiddleware.AuthMiddleware()).GET("/users/:uuid", handler.GetUserByUUID)
	r.Use(authMiddleware.AuthMiddleware()).GET("/user", handler.GetUser)
}
