package handler

import (
	"sk-go-be/internal/model"
	"sk-go-be/internal/service"
)

type PostHandler struct {
	PostService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{PostService: postService}
}

func (h *PostHandler) GetPostByUUID(uuid string) (*model.Post, error) {
	// Parse the UUID
	// parsedUUID, err := uuid.Parse(uuid)
	// if err != nil {
	// 	return nil, err
	// }

	// return s.repo.GetPostByUUID(parsedUUID)

	return nil, nil
}
