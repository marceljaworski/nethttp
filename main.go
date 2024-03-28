package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", func(w http.ResponseWriter, r *http.Request) {

	})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
