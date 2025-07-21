package repository

import (
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProductByUUID(uuid string) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of productRepository.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductByUUID(uuid string) (*model.Product, error) {
	var product model.Product
	err := r.db.Where("uuid = ?", uuid).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No product found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &product, nil
}
