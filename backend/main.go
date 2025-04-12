package main

import (
	"forum/db"
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

	listPostHandler := InitializeListPostHandler(db)
	getPostHandler := InitializeGetPostHandler(db)
	createPostHandler := InitializeCreatePostHandler(db)
	deletePostHandler := InitializeDeletePostHandler(db)

	router := http.NewServeMux()
	router.Handle("GET /posts", listPostHandler)
	router.Handle("GET /posts/{id}", getPostHandler)
	router.Handle("POST /posts", createPostHandler)
	router.Handle("DELETE /posts/{id}", deletePostHandler)

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", router))

	server := &http.Server{
		Addr:    ":4000",
		Handler: api,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}
}
