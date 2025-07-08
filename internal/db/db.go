package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Bangkok",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected!")

	RunMigrations(db)

	return db
}

// runMigrations applies database migrations using golang-migrate
func RunMigrations(database *gorm.DB) error {
	// Open a raw database connection for golang-migrate
	log.Println("Start Applying migrations")
	db, err := database.DB()
	if err != nil {
		log.Fatalf("failed to get raw DB connection: %v", err)
		return err
	}

	// Initialize the postgres driver for golang-migrate
	driver, err := migratePostgres.WithInstance(db, &migratePostgres.Config{})
	if err != nil {
		log.Fatalf("failed to initialize postgres driver: %v", err)
		return err
	}

	// Initialize migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Path to migration files
		"postgres",          // Database name (used by golang-migrate)
		driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize migrate: %v", err)
		return err
	}

	// Apply migrations
	log.Println("Applying migrations...")
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply")
		} else {
			log.Fatalf("failed to apply migrations: %v", err)
			return err
		}
	}

	// Ensure the database connection is still valid
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database after migrations: %v", err)
		return err
	}

	return nil
}
