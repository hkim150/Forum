package service

import (
	"context"
	"forum/internal/model"
	"forum/internal/repository"
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService(repository *repository.PostRepository) *PostService {
	return &PostService{
		repository: repository,
	}
}

func (s *PostService) List(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	posts, err := s.repository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) Get(ctx context.Context, id int) (*model.Post, error) {
	return s.repository.Get(ctx, id)
}

func (s *PostService) Create(ctx context.Context, title, content string) (int, error) {
	return s.repository.Create(ctx, repository.PostRepositoryCreateParams{
		Title:   title,
		Content: content,
	})
}

func (s *PostService) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
