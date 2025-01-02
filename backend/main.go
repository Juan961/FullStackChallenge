package main

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set up CORS middleware with custom configuration
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"X-Custom-Header"},
		AllowCredentials: true,
		MaxAge:           300, // Cache the CORS preflight response for 5 minutes
	}).Handler)

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		response := List()
		
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		query := queryParams.Get("query")

		results := Search(query)
		
		w.Header().Set("Content-Type", "application/json")
		w.Write(results)
	})

	http.ListenAndServe(":3333", r)
}
