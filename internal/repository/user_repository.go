package repository

import (
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

// UserRepository handles user data access

type UserRepository interface {
	GetUserByUUID(uuid string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUUID(uuid string) (*model.User, error) {
	var user model.User
	err := r.db.Where("uuid = ?", uuid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No user found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &user, nil
}
