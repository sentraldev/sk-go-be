package repository

import (
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetProductCategoryList() ([]model.Category, error)
	CreateNewProductCategory(category *model.Category) error
	UpdateProductCategory(category *model.Category) error
	DeleteProductCategory(uuid string) error // Optional: If you want to support deletion
}

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new instance of categoryRepository.
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetProductCategoryList() ([]model.Category, error) {
	var categories []model.Category
	// Sort by name in ascending order
	if err := r.db.Order("name ASC").Find(&categories).Error; err != nil {
		return nil, err // Handle error if query fails
	}
	return categories, nil
}

func (r *categoryRepository) CreateNewProductCategory(category *model.Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err // Handle error if creation fails
	}
	return nil // Successfully created the category
}

func (r *categoryRepository) UpdateProductCategory(category *model.Category) error {
	existingCategory := &model.Category{}
	if err := r.db.First(existingCategory, "uuid = ?", category.UUID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound // Category not found
		}
		return err // Other errors
	}
	existingCategory.Name = category.Name

	if err := r.db.Save(existingCategory).Error; err != nil {
		return err // Handle error if update fails
	}
	return nil // Successfully updated the category
}

func (r *categoryRepository) DeleteProductCategory(uuid string) error {
	if err := r.db.Delete(&model.Category{}, "uuid = ?", uuid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound // Category not found
		}
		return err // Other errors
	}
	return nil // Successfully deleted the category
}
