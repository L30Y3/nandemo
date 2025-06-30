package main

import (
	"log"
	"net/http"
	"os"

	"github.com/L30Y3/nandemo/merchant-service/internal/consumer"
	"github.com/L30Y3/nandemo/shared/events"
)

func main() {
	bus := events.NewInMemoryBus()
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
