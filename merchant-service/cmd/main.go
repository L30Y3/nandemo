package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/L30Y3/nandemo/merchant-service/internal/api"
	"github.com/L30Y3/nandemo/merchant-service/internal/consumer"
	cfg "github.com/L30Y3/nandemo/shared/config"
	"github.com/L30Y3/nandemo/shared/events"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Define command line args, run like below or run without args to use defaults, add to README
	// 	go run main.go \
	//   -project=my-alt-project \
	//   -topic=my-alt-topic \
	//   -subscription=my-alt-sub

	projectID := flag.String("project", cfg.DefaultProjectID, "GCP project ID")
	topicID := flag.String("topic", cfg.DefaultTopicID, "Pub/Sub topic ID")
	subID := flag.String("subscription", cfg.DefaultSubID, "Pub/Sub subscription ID")

	flag.Parse()

	ctx := context.Background()
	var bus *events.PubSubBus
	var err error

	if os.Getenv("DISABLE_PUBSUB") != "true" {
		log.Println("Pub/Sub is enabled, using real bus")
		bus, err = events.NewPubSubSubscriber(ctx, *projectID, *topicID, *subID)
		if err != nil {
			log.Fatalf("Failed to create PubSubBus: %v", err)
		}

		defer bus.Stop()
	}

	client, err := firestore.NewClient(ctx, cfg.DefaultProjectID)
	if err != nil {
		log.Fatalf("Failed to init Firestore client: %v", err)
	}

	defer client.Close()

	if bus != nil {
		go consumer.ListenForOrders(bus, client)
	}

	r := chi.NewRouter()
	api.RegisterRoutes(r, &api.MerchantHandlerWithFirestoreClient{
		Firestore: client,
	})

	var port string
	if os.Getenv("IN_CONTAINER") == "true" {
		port = "80"
	} else {
		port = cfg.MerchantServicePort
		log.Printf("Using default port: %s", port)
	}

	log.Printf("Merchant Service running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
