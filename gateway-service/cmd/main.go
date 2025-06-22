package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/gateway-service/internal/api"
)

func main() {
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Gateway Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
