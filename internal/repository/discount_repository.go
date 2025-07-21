package repository

import "gorm.io/gorm"

type DiscountRepository interface {
	// Add methods for discount repository here
}

type discountRepository struct {
	// db *gorm.DB // Uncomment if using GORM for database operations
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepository{
		db: db,
	}
}

// Implement methods for DiscountRepository here
