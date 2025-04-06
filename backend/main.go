package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", handler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Hello, World from Go Server!</h1>"))
}
