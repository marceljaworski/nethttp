package main

import (
	"log"
	"net/http"
)

func findByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("received request for item: " + id))
}
func main() {
	router := http.NewServeMux()
	// Path parameter id
	router.HandleFunc("/item/{id}", findByID)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
