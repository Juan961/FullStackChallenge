package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Set environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Set PORT
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "3333"
	}

	// Set credentials for ZincSearch
	user := os.Getenv("ZINCSEARCH_ADMIN_USER")
	password := os.Getenv("ZINCSEARCH_ADMIN_PASSWORD")
	if user == "" || password == "" {
		fmt.Println("Environment variables SEARCH_USER and SEARCH_PASSWORD must be set")
		return
	}
	creds := user + ":" + password
	bas64encodedCreds := base64.StdEncoding.EncodeToString([]byte(creds))

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set up CORS
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"X-Custom-Header"},
		AllowCredentials: true,
		MaxAge:           300, // Cache the CORS preflight response for 5 minutes
	}).Handler)

	// Set index info endpoint
	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		statusCode, error, response := List(string(bas64encodedCreds))

		w.WriteHeader(statusCode)

		if error != nil {
			w.Write([]byte(error.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	// Set search endpoint
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		query := queryParams.Get("query")

		statusCode, error, results := Search(string(bas64encodedCreds), query)

		w.WriteHeader(statusCode)

		if error != nil {
			w.Write([]byte(error.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(results)
	})

	http.ListenAndServe(":"+port, r)
}
