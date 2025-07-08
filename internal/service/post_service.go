package service

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/repository"
)

type PostService interface {
	GetPostByUUID(uuid string) (*model.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) GetPostByUUID(uuid string) (*model.Post, error) {
	// Parse the UUID
	// parsedUUID, err := uuid.Parse(uuid)
	// if err != nil {
	// 	return nil, err
	// }

	// return s.repo.GetPostByUUID(parsedUUID)

	return nil, nil
}
