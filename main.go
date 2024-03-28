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

// func handlerAdmin(w http.ResponseWriter, r *http.Request) {
// userID, ok := r.Context().Value(middleware.AuthUserID).(string)
// if !ok {
// 	log.Println("invalid user ID")
// 	w.WriteHeader(http.StatusBadRequest)
// }
// w.Write([]byte("you are an authorized Admin " + userID))
// }

func main() {
	router := http.NewServeMux()

	// specific method GET and Path parameter id
	router.HandleFunc("GET /item/{id}", findByID)

	// Subrouting
	router.Handle("/v1/", http.StripPrefix("/v1", router))

	// Router restricted to Admins only
	// adminRouter := http.NewServeMux()
	// adminRouter.HandleFunc("POST /item/", handlerAdmin)
	// adminRouter.HandleFunc("PUT /item/{id}", handlerAdmin)
	// adminRouter.HandleFunc("DELETE /item/{id}", handlerAdmin)
	// Middleware for Admin Authentication
	// router.Handle("/", middleware.EnsureAdmin(adminRouter))

	stack := middleware.CreateStack(
		middleware.Logging,
		// middleware.EnsureAdmin
		// middleware.AllowCors,
		// middleware.IsAuthentiticated,
		// middleware.CheckPerissions,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
