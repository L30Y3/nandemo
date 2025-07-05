package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/gateway-service/internal/api"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	api.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Gateway Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
