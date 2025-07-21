package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

// UserService contains business logic for users

type UserService interface {
	GetUserByExteralID(user_id string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByExteralID(user_id string) (*model.User, error) {
	return s.repo.GetUserByExteralID(user_id)
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, error) {
	// Parse the UUID
	// parsedUUID, err := uuid.Parse(uuid)
	// if err != nil {
	// 	return nil, err
	// }

	// return s.repo.GetUserByUUID(parsedUUID)

	return nil, nil
}
