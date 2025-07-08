package repository

import (
	"sk-go-be/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetShoppingCartByUUID(uuid uuid.UUID) (*model.Cart, error)
	GetShoppingCartByUserUUID(userUUID uuid.UUID) (*model.Cart, error)
}

type cartRepository struct {
	db *gorm.DB
}

// NewCartRepository creates a new instance of cartRepository.
func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetShoppingCartByUUID(uuid uuid.UUID) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Where("uuid = ?", uuid).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No cart found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &cart, nil
}

func (r *cartRepository) GetShoppingCartByUserUUID(userUUID uuid.UUID) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Where("user_uuid = ?", userUUID).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No cart found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &cart, nil
}
