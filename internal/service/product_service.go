package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

type ProductService interface {
	GetProducts() ([]model.Product, error)
	GetProductByUUID(uuid string) (*model.Product, error)

	// ADMIN
	CreateNewProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(uuid string) error // Optional: If you want to support deletion
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProducts() ([]model.Product, error) {
	products, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) GetProductByUUID(uuid string) (*model.Product, error) {
	// Parse the UUID
	// parsedUUID, err := uuid.Parse(uuid)
	// if err != nil {
	// 	return nil, err
	// }

	return s.repo.GetProductByUUID(uuid)

	// return nil, nil
}

// ADMIN
func (s *productService) CreateNewProduct(product *model.Product) error {
	if err := s.repo.CreateNewProduct(product); err != nil {
		return err // Handle error if creation fails
	}
	return nil // Successfully created the product
}

func (s *productService) UpdateProduct(product *model.Product) error {
	if err := s.repo.UpdateProduct(product); err != nil {
		return err // Handle error if update fails
	}
	return nil // Successfully updated the product
}

func (s *productService) DeleteProduct(uuid string) error {
	if err := s.repo.DeleteProduct(uuid); err != nil {
		return err // Handle error if deletion fails
	}
	return nil // Successfully deleted the product
}
