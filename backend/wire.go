//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"forum/handler/post"
	"forum/repository"
	"forum/service"

	"github.com/google/wire"
)

func InitializeListPostHandler(db *sql.DB) *post.ListPostHandler {
	wire.Build(
		post.NewListPostHandler,
		service.NewPostService,
		repository.NewPostRepository,
	)
	return &post.ListPostHandler{}
}

func InitializeGetPostHandler(db *sql.DB) *post.GetPostHandler {
	wire.Build(
		post.NewGetPostHandler,
		service.NewPostService,
		repository.NewPostRepository,
	)
	return &post.GetPostHandler{}
}

func InitializeCreatePostHandler(db *sql.DB) *post.CreatePostHandler {
	wire.Build(
		post.NewCreatePostHandler,
		service.NewPostService,
		repository.NewPostRepository,
	)
	return &post.CreatePostHandler{}
}

func InitializeDeletePostHandler(db *sql.DB) *post.DeletePostHandler {
	wire.Build(
		post.NewDeletePostHandler,
		service.NewPostService,
		repository.NewPostRepository,
	)
	return &post.DeletePostHandler{}
}
