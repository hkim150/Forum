package server

import (
	"fmt"
	"forum/internal/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func Run(port string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	db, err := db.Open()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return
	}
	defer db.Close()

	listPostHandler := InitializeListPostHandler(db)
	getPostHandler := InitializeGetPostHandler(db)
	createPostHandler := InitializeCreatePostHandler(db)
	deletePostHandler := InitializeDeletePostHandler(db)

	router := http.NewServeMux()
	router.Handle("GET /posts", listPostHandler)
	router.Handle("GET /posts/{id}", getPostHandler)
	router.Handle("POST /posts", createPostHandler)
	router.Handle("DELETE /posts/{id}", deletePostHandler)

	subPath := http.NewServeMux()
	subPath.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: subPath,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}
}
