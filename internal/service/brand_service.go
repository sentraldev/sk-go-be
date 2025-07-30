package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

type BrandService interface {
	GetBrands() ([]model.Brand, error)
	GetBrandByUUID(uuid string) (*model.Brand, error)
	CreateNewBrand(brand *model.Brand) error
	UpdateBrand(brand *model.Brand) error
	DeleteBrand(uuid string) error // Optional: If you want to support deletion
}

type brandService struct {
	repo repository.BrandRepository
}

func NewBrandService(repo repository.BrandRepository) BrandService {
	return &brandService{repo: repo}
}

func (s *brandService) GetBrands() ([]model.Brand, error) {
	brands, err := s.repo.GetBrands()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (s *brandService) GetBrandByUUID(uuid string) (*model.Brand, error) {
	brand, err := s.repo.GetBrandByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (s *brandService) CreateNewBrand(brand *model.Brand) error {
	return s.repo.CreateNewBrand(brand)
}

func (s *brandService) UpdateBrand(brand *model.Brand) error {
	return s.repo.UpdateBrand(brand)
}

func (s *brandService) DeleteBrand(uuid string) error {
	return s.repo.DeleteBrand(uuid)
}
