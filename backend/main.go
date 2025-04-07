package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func main() {
	http.HandleFunc("/api/data", handler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
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