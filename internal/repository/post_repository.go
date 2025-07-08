package repository

import (
	"sk-go-be/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetPostByUUID(uuid string) (*model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

// NewPostRepository creates a new instance of postRepository.
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetPostByUUID(uuid string) (*model.Post, error) {
	var post model.Post
	err := r.db.Where("uuid = ?", uuid).First(&post).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // No post found
	}
	if err != nil {
		return nil, err // Other database errors
	}
	return &post, nil
}
