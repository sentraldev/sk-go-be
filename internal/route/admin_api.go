package route

import (
	"sk-go-be/internal/handler"
	"sk-go-be/internal/handler/admin"

	"sk-go-be/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminAPIRoutes(
	r *gin.Engine,
	fc *middleware.FirebaseClient,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	wishlistHandler *handler.WishlistHandler,
	discountHandler *handler.DiscountHandler,
	authHandler *handler.AuthHandler,
	categoryHandler *admin.CategoryHandler,
) {
	v1 := r.Group("/api/v1")

	// Admin API routes
	admin := v1.Group("/admin")
	admin.Use(fc.AuthMiddleware()) // Ensure admin authentication
	// Add admin endpoints here
	{
		// Category routes
		SetupAdminCategoryRoutes(admin, *categoryHandler)

		// admin.GET("/products", handler.ListUsers)
		// SetupAdminProductRoutes(admin, *productHandler)

		// SetupAdminCartRoutes(admin, *cartHandler)
		// SetupAuthRoutes(admin, *authHandler, *fc)
		// SetupUserRoutes(admin, *userHandler, *fc)
	}
}

func SetupAdminCategoryRoutes(r *gin.RouterGroup, handler admin.CategoryHandler) {
	r.GET("/categories", handler.GetProductCategoryList)
	r.POST("/categories", handler.CreateNewProductCategory)
	r.POST("/categories/:uuid", handler.UpdateProductCategory) // Ensure this is the correct method for creating categories
	// Optional: If you want to support deletion
	// Proceed with caution as this will delete the category
	r.POST("/categories/delete/:uuid", handler.DeleteProductCategory)
}

// func SetupAdminCartRoutes(r *gin.RouterGroup, handler handler.CartHandler) {
// 	r.GET("/cart/:uuid", handler.GetShoppingCartByUUID)
// 	r.GET("/cart/user/:uuid", handler.GetShoppingCartByUserUUID)
// }

// func SetupAdminProductRoutes(r *gin.RouterGroup, handler handler.ProductHandler) {
// 	r.GET("/products/:uuid", handler.GetProductByUUID)
// 	r.GET("/products", handler.GetProducts)
// }
