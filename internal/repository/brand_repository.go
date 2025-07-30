package repository

import (
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

type BrandRepository interface {
	GetBrands() ([]model.Brand, error)
	GetBrandByUUID(uuid string) (*model.Brand, error)
	CreateNewBrand(brand *model.Brand) error
	UpdateBrand(brand *model.Brand) error
	DeleteBrand(uuid string) error // Optional: If you want to support deletion
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}
func (r *brandRepository) GetBrands() ([]model.Brand, error) {
	var brands []model.Brand
	err := r.db.Order("name ASC").Find(&brands).Error
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (r *brandRepository) GetBrandByUUID(uuid string) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.Where("uuid = ?", uuid).First(&brand).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No brand found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &brand, nil
}

func (r *brandRepository) CreateNewBrand(brand *model.Brand) error {
	if err := r.db.Create(brand).Error; err != nil {
		return err // Handle error if creation fails
	}
	return nil // Successfully created the brand
}

func (r *brandRepository) UpdateBrand(brand *model.Brand) error {
	if err := r.db.Save(brand).Error; err != nil {
		return err // Handle error if update fails
	}
	return nil // Successfully updated the brand
}

func (r *brandRepository) DeleteBrand(uuid string) error {
	if err := r.db.Delete(&model.Brand{}, "uuid = ?", uuid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound // Brand not found
		}
		return err // Other errors
	}
	return nil // Successfully deleted the brand
}
