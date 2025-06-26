package main

import (
	"log"

	"github.com/L30Y3/nandemo/merchant-service/internal/consumer"
	"github.com/L30Y3/nandemo/shared/events"
)

func main() {
	bus := events.NewInMemoryBus()

	// Simulate event subscription (in-memory)
	go consumer.ListenForOrders(bus)

	log.Println("Merchant Service running...")

	// Block forever
	select {}
}
