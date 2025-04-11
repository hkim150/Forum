package service

import (
	"context"
	"forum/repository"
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService(repository *repository.PostRepository) *PostService {
	return &PostService{
		repository: repository,
	}
}

func (s *PostService) Create(ctx context.Context, title, content string) (int, error) {
	return s.repository.Create(ctx, repository.PostRepositoryCreateParams{
		Title:   title,
		Content: content,
	})
}
