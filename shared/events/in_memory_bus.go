package events

import (
	"fmt"

	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

type InMemoryBus struct {
	orderCreatedSubscribers []OrderCreatedEventHandler
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{
		orderCreatedSubscribers: []OrderCreatedEventHandler,
	}
}

func (b *InMemoryBus) PublishOrderCreated(event *pb.OrderCreatedEvent) error {
	fmt.Printf("In-Memory Event Published: %+v\n", event)

	for _, handler := range b.orderCreatedSubscribers {
		// may make this async later
		handler(event)
	}
	return nil
}

func (b *InMemoryBus) SubscribeToOrderCreated(handler OrderCreatedEventHandler) {
	b.orderCreatedSubscribers = append(b.orderCreatedSubscribers, handler)
	fmt.Println("Subscriber registered for OrderCreatedEvent.")
}
