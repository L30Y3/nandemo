package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/order-service/internal/api"
	"github.com/L30Y3/nandemo/shared/events"
)

func main() {
	mux := http.NewServeMux()

	bus := events.NewInMemoryBus()
	api.RegisterRoutes(mux, bus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Starting Order Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
