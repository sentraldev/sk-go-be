package repository

import (
	"sk-go-be/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProductByUUID(uuid string) (*model.Product, error)

	// ADMIN
	CreateNewProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(uuid string) error // Optional: If you want to support deletion
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
	err := r.db.Preload("Category").Preload("Brand").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductByUUID(productUuid string) (*model.Product, error) {
	var product model.Product
	parsedUuid, err := uuid.Parse(productUuid)
	if err != nil {
		return nil, err
	}
	err = r.db.Preload("Category").Preload("Brand").Where("uuid = ?", parsedUuid).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No product found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &product, nil
}

// ADMIN
func (r *productRepository) CreateNewProduct(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err // Handle error if creation fails
	}
	return nil // Successfully created the product
}

func (r *productRepository) UpdateProduct(product *model.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err // Handle error if update fails
	}
	return nil // Successfully updated the product
}

func (r *productRepository) DeleteProduct(uuid string) error {
	if err := r.db.Delete(&model.Product{}, "uuid = ?", uuid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound // Product not found
		}
		return err // Other errors
	}
	return nil // Successfully deleted the product
}
