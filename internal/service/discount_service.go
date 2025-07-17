package service

import (
	"sk-go-be/internal/repository"

	"github.com/google/uuid"
)

type DiscountService interface {
	// Define methods for discount service here
}

type discountService struct {
	repo repository.DiscountRepository
}

func NewDiscountService(repo repository.DiscountRepository) DiscountService {
	return &discountService{
		repo: repo,
	}
}

func (s *discountService) AttachDiscountToProduct(productUUID uuid.UUID, percentage float32) error {
	// TODO: Implement the logic to attach a discount to a product

	// TODO: Check if the product exists

	// TODO: Check if the discount already exists for the product

	// TODO: If the discount does not exist, create a new discount.

	// TODO: If the discount already exists, update the existing discount.

	// TODO: Attach the discount to the product

	return nil
}

func (s *discountService) RemoveDiscountFromProduct(productUUID uuid.UUID) error {
	// TODO: Implement the logic to remove a discount from a product

	// TODO: Check if the product exists

	// TODO: Check if the discount exists for the product

	// TODO: If the discount exists, remove it. Else, return an error indicating that no discount is attached to the product.

	return nil
}
