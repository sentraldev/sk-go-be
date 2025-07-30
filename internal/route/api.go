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
	categoryHandler *handler.CategoryHandler,
	brandHandler *handler.BrandHandler,
) {
	v1 := r.Group("/api/v1")

	// Admin API routes
	adminGroup := v1.Group("/admin")
	adminGroup.Use(fc.AuthMiddleware()) // Admin authentication
	SetupAdminCategoryRoutes(adminGroup, *categoryHandler)
	SetupAdminProductRoutes(adminGroup, *productHandler)
	SetupAdminBrandRoutes(adminGroup, *brandHandler)

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
}

func SetupCartPublicRoutes(r *gin.RouterGroup, handler handler.CartHandler) {
	r.GET("/cart/:uuid", handler.GetShoppingCartByUUID)
	r.GET("/cart/user/:uuid", handler.GetShoppingCartByUserUUID)
}

func SetupProductRoutes(r *gin.RouterGroup, handler handler.ProductHandler) {
	r.GET("/products/:uuid", handler.GetProductByUUID)
	r.GET("/products", handler.GetProducts)
}

func SetupAuthRoutes(r *gin.RouterGroup, handler handler.AuthHandler, authMiddleware middleware.FirebaseClient) {
	r.Use(authMiddleware.AuthMiddleware()).POST("/register", handler.Register)
	// r.POST("/login", handler.Login) // Uncomment when login is implemented
}

func SetupUserRoutes(r *gin.RouterGroup, handler handler.UserHandler, authMiddleware middleware.FirebaseClient) {
	r.Use(authMiddleware.AuthMiddleware()).GET("/users/:uuid", handler.GetUserByUUID)
	r.Use(authMiddleware.AuthMiddleware()).GET("/user", handler.GetUser)
}

/// ADMIN

// Admin endpoints
func SetupAdminCategoryRoutes(r *gin.RouterGroup, handler handler.CategoryHandler) {
	r.GET("/categories", handler.GetProductCategoryList)
	r.POST("/categories", handler.CreateNewProductCategory)
	r.POST("/categories/:uuid", handler.UpdateProductCategory)
	r.POST("/categories/delete/:uuid", handler.DeleteProductCategory) // Use with caution
}

func SetupAdminProductRoutes(r *gin.RouterGroup, handler handler.ProductHandler) {
	// r.GET("/products/:uuid", handler.GetProductByUUID)
	r.GET("/products", handler.GetProducts)
	r.POST("/products", handler.CreateNewProduct)
}

func SetupAdminBrandRoutes(r *gin.RouterGroup, handler handler.BrandHandler) {
	// r.GET("/brands/:uuid", handler.GetBrandByUUID)
	r.GET("/brands", handler.GetBrands)
	r.POST("/brands", handler.CreateNewBrand)
	r.POST("/brands/:uuid", handler.UpdateBrand)
	r.POST("/brands/delete/:uuid", handler.DeleteBrand) // Use with caution
}
