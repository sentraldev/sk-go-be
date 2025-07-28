package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

type CategoryService interface {
	GetProductCategoryList() ([]model.Category, error)
	CreateNewProductCategory(category *model.Category) error
	UpdateProductCategory(category *model.Category) error
	DeleteProductCategory(uuid string) error // Optional: If you want to support deletion
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetProductCategoryList() ([]model.Category, error) {
	categories, err := s.repo.GetProductCategoryList()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *categoryService) CreateNewProductCategory(category *model.Category) error {
	if err := s.repo.CreateNewProductCategory(category); err != nil {
		return err
	}
	return nil
}

func (s *categoryService) UpdateProductCategory(category *model.Category) error {
	if err := s.repo.UpdateProductCategory(category); err != nil {
		return err
	}
	return nil
}

func (s *categoryService) DeleteProductCategory(uuid string) error {
	if err := s.repo.DeleteProductCategory(uuid); err != nil {
		return err
	}
	return nil
}
