package main

import (
	"log"
	"net/http"

	"github.com/marceljaworski/nethttp/middleware"
)

func findByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	method := r.Method
	w.Write([]byte("received " + method + " request for item: " + id))
}

func main() {
	router := http.NewServeMux()
	// Path parameter id
	router.HandleFunc("GET /item/{id}", findByID)
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
