package repository

import (
	"log"
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

// UserRepository handles user data access

type UserRepository interface {
	GetUserByExteralID(externalID string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, error)
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByExteralID(externalID string) (*model.User, error) {
	log.Printf("User UID from context: %s", externalID)

	var user model.User
	var users []model.User

	if err1 := r.db.Find(&users).Error; err1 != nil {
		return nil, err1
	}

	log.Printf("Found %d users in the database", len(users))

	err := r.db.Where("external_user_id = ?", externalID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No user found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &user, nil
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

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
