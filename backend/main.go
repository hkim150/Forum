package main

import (
	"encoding/json"
	"forum/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	db, err := db.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/api/data", handler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := "TODO: add data struct"

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
