package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"sk-go-be/internal/db"
	"sk-go-be/internal/handler"
	"sk-go-be/internal/middleware"
	"sk-go-be/internal/repository"
	"sk-go-be/internal/route"
	"sk-go-be/internal/service"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func main() {
	// cfg := config.Load()
	dbConn := db.ConnectDatabase()
	r := gin.Default()

	var opt option.ClientOption
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
		opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS"))) // Load from environment variable
	} else {
		gin.SetMode(gin.DebugMode)
		// Check if file not found
		if _, err := os.Stat("sk-web-2025-firebase.json"); os.IsNotExist(err) {
			opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS"))) // Load from environment variable
		} else {
			opt = option.WithCredentialsFile("sk-web-2025-firebase.json")
		}
	}

	ctx := context.Background()

	fbApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}

	authClient, err := fbApp.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth: %v", err)
	}

	fc := middleware.NewFirebaseClient(authClient)

	userHandler, productHandler, cartHandler, wishlistHandler, discountHandler, authHandler := SetupHandler(dbConn, fbApp)

	allowedOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	originsList := strings.Split(allowedOrigins, ",") // Convert string to slice
	log.Println(originsList)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     originsList,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.RedirectTrailingSlash = true
	route.RegisterAPIRoutes(r, *fc,
		userHandler,
		productHandler,
		cartHandler,
		wishlistHandler,
		discountHandler,
		authHandler,
	)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func SetupHandler(db *gorm.DB, fbApp *firebase.App) (
	*handler.UserHandler,
	*handler.ProductHandler,
	*handler.CartHandler,
	*handler.WishlistHandler,
	*handler.DiscountHandler,
	*handler.AuthHandler,
) {
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	authService := service.NewAuthService(userRepo, fbApp)
	authHandler := handler.NewAuthHandler(authService)

	// Shopping Cart
	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	// Wishlist
	wishlistRepo := repository.NewWishlistRepository(db)
	wishlistService := service.NewWishlistService(wishlistRepo)
	wishlistHandler := handler.NewWishlistHandler(wishlistService)

	// Discount
	discountRepo := repository.NewDiscountRepository(db)
	discountService := service.NewDiscountService(discountRepo)
	discountHandler := handler.NewDiscountHandler(discountService)
	return userHandler, productHandler, cartHandler, wishlistHandler, discountHandler, authHandler

}
