package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/gateway-service/internal/api"
	"github.com/L30Y3/nandemo/shared/clients/merchantclient"
	"github.com/L30Y3/nandemo/shared/clients/orderclient"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Initialize the API handler
	handler := &api.Handler{
		OrderClient:    orderclient.NewOrderServiceClient(),
		MerchantClient: merchantclient.NewMerchantServiceClient(),
	}

	// Register routes with the handler
	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // change to specific frontend domain in prod
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Gateway Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
