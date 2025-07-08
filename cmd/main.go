package main

import (
	"log"

	"sk-go-be/internal/db"
	"sk-go-be/internal/handler"
	"sk-go-be/internal/repository"
	"sk-go-be/internal/route"
	"sk-go-be/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// cfg := config.Load()
	dbConn := db.ConnectDatabase()
	r := gin.Default()

	userHandler, productHandler, cartHandler, wishlistHandler := SetupHandler(dbConn)

	route.RegisterAPIRoutes(r,
		userHandler,
		productHandler,
		cartHandler,
		wishlistHandler,
	)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func SetupHandler(db *gorm.DB) (
	*handler.UserHandler,
	*handler.ProductHandler,
	*handler.CartHandler,
	*handler.WishlistHandler,
) {
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Shopping Cart
	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	// Wishlist
	wishlistRepo := repository.NewWishlistRepository(db)
	wishlistService := service.NewWishlistService(wishlistRepo)
	wishlistHandler := handler.NewWishlistHandler(wishlistService)

	return userHandler, productHandler, cartHandler, wishlistHandler
}
