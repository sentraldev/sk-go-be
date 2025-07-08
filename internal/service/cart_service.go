package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"

	"github.com/google/uuid"
)

type CartService interface {
	GetShoppingCartByUUID(uuid string) (*model.Cart, error)
	GetShoppingCartByUserUUID(userUUID string) (*model.Cart, error)
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}
}

func (s *cartService) GetShoppingCartByUUID(cartUuid string) (*model.Cart, error) {
	// Parse the UUID
	parsedUUID, err := uuid.Parse(cartUuid)
	if err != nil {
		return nil, err
	}

	return s.repo.GetShoppingCartByUUID(parsedUUID)
}

func (s *cartService) GetShoppingCartByUserUUID(userUUID string) (*model.Cart, error) {
	// Parse the UUID
	parsedUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, err
	}

	return s.repo.GetShoppingCartByUserUUID(parsedUUID)
}
