package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../frontend/static"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
