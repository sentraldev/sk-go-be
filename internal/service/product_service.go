package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

type ProductService interface {
	GetProducts() ([]model.Product, error)
	GetProductByUUID(uuid string) (*model.Product, error)
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

	// return s.repo.GetProductByUUID(parsedUUID)

	return nil, nil
}
