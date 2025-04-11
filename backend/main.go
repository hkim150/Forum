package main

import (
	"forum/db"
	"forum/handler/post"
	"forum/repository"
	"forum/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	db, err := db.NewDb()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return
	}
	defer db.Close()

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	createPostHandler := post.NewCreatePostHandler(postService)

	router := http.NewServeMux()
	router.Handle("POST /api/posts", createPostHandler)

	server := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}
}
