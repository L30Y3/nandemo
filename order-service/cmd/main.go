package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/order-service/internal/api"
	"github.com/L30Y3/nandemo/shared/events"
)

const (
	defaultProjectID = "nandemo-464411"
	defaultTopicID   = "order-created"
)

func main() {
	mux := http.NewServeMux()

	ctx := context.Background()

	// this is a publishing only service, don't need a real subscription ID
	bus, err := events.NewPubSubBus(ctx, defaultProjectID, defaultTopicID, "not-used-in-publisher-mode")
	if err != nil {
		log.Fatalf("Failed to create PubSubBus: %v", err)
	}

	defer bus.Stop()
	api.RegisterRoutes(mux, bus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Starting Order Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
