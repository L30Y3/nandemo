package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/merchant-service/internal/consumer"
	"github.com/L30Y3/nandemo/shared/events"
)

const (
	defaultProjectID      = "nandemo-464411"
	defaultTopicID        = "order-created"
	defaultSubscriptionID = "merchant-order-receiver"
)

func main() {
	// Define command line args, run like below or run without args to use defaults, add to README
	// 	go run main.go \
	//   -project=my-alt-project \
	//   -topic=my-alt-topic \
	//   -subscription=my-alt-sub

	projectID := flag.String("project", defaultProjectID, "GCP project ID")
	topicID := flag.String("topic", defaultTopicID, "Pub/Sub topic ID")
	subID := flag.String("subscription", defaultSubscriptionID, "Pub/Sub subscription ID")

	flag.Parse()

	ctx := context.Background()
	bus, err := events.NewPubSubBus(ctx, *projectID, *topicID, *subID)
	if err != nil {
		log.Fatalf("Failed to create PubSubBus: %v", err)
	}

	defer bus.Stop()

	consumer.ListenForOrders(bus)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Merchant Service OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Merchant Service running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
