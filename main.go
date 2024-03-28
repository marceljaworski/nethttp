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

	// specific method GET and Path parameter id
	router.HandleFunc("GET /item/{id}", findByID)

	// Subrouting
	// v1 := http.NewServeMux()
	router.Handle("/v1/", http.StripPrefix("/v1", router))

	stack := middleware.CreateStack(
		middleware.Logging,
		// middleware.AllowCors,
		// middleware.IsAuthed,
		// middleware.CheckPerissions,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
