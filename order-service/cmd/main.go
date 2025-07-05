package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/order-service/internal/api"
	cfg "github.com/L30Y3/nandemo/shared/config"
	"github.com/L30Y3/nandemo/shared/events"
)

func main() {
	mux := http.NewServeMux()

	ctx := context.Background()

	// this is a publishing only service, don't need a real subscription ID
	bus, err := events.NewPubSubPublisher(ctx, cfg.DefaultProjectID, cfg.DefaultTopicID)
	if err != nil {
		log.Fatalf("Failed to create PubSubBus: %v", err)
	}

	defer bus.Stop()
	api.RegisterRoutes(mux, bus)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.OrderServicePort
	}

	log.Printf("Starting Order Service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
