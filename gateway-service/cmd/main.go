package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/gateway-service/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // change to specific frontend domain in prod
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	api.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Gateway Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
