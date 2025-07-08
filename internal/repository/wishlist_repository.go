package repository

import (
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

type WishlistRepository interface {
	GetWishlistByUserUUID(uuid string) (*model.Wishlist, error)
}

type wishlistRepository struct {
	db *gorm.DB
}

// NewWishlistRepository creates a new instance of wishlistRepository.
func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{db: db}
}

func (r *wishlistRepository) GetWishlistByUserUUID(uuid string) (*model.Wishlist, error) {
	var wishlist model.Wishlist
	err := r.db.Where("user_uuid = ?", uuid).First(&wishlist).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No wishlist found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &wishlist, nil
}
