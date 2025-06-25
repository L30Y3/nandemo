package events

import (
	"fmt"

	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

type InMemoryBus struct{}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{}
}

func (b *InMemoryBus) PublishOrderCreated(event *pb.OrderCreatedEvent) error {
	fmt.Printf("In-Memory Event Published: %+v\n", event)
	return nil
}
