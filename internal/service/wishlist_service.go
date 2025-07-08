package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

type WishlistService interface {
	GetWishlistByUserUUID(uuid string) (*model.Wishlist, error)
}

type wishlistService struct {
	repo repository.WishlistRepository
}

func NewWishlistService(repo repository.WishlistRepository) WishlistService {
	return &wishlistService{repo: repo}
}

func (s *wishlistService) GetWishlistByUserUUID(uuid string) (*model.Wishlist, error) {
	// Parse the UUID
	// parsedUUID, err := uuid.Parse(uuid)
	// if err != nil {
	// 	return nil, err
	// }

	// return s.repo.GetWishlistByUserUUID(parsedUUID)

	return nil, nil
}
