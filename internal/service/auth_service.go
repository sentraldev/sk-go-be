package service

import (
	"errors"
	"log"
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"

	firebase "firebase.google.com/go/v4"
	"github.com/google/uuid"
)

type AuthService interface {
	// Login(email, password string) (string, error)
	Register(uid, name, phone, email string) error
}

type authService struct {
	userRepo repository.UserRepository
	fbApp    *firebase.App
}

func NewAuthService(userRepo repository.UserRepository, fbApp *firebase.App) AuthService {
	return &authService{userRepo: userRepo, fbApp: fbApp}
}

func (s *authService) Login(email, password string) (string, error) {
	// Firebase Admin SDK does not support password login directly.
	// Use Firebase REST API from frontend for password authentication.
	return "", errors.New("email/password login must be handled via Firebase REST API from frontend")
}

func (s *authService) Register(uid, name, phone, email string) error {
	user := &model.User{
		UUID:           uuid.New(),
		ExternalUserID: uid,
		Name:           name,
		Phone:          phone,
		Email:          email,
		Role:           "user", // Default role, can be changed later
	}

	log.Println("Registering user:", user)

	return s.userRepo.CreateUser(user)
}
