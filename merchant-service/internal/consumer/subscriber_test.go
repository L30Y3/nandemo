package consumer

import (
	"testing"

	"github.com/L30Y3/nandemo/shared/events"
	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

func TestOrderEventDeliveryToMerchant(t *testing.T) {
	bus := events.NewInMemoryBus()

	received := make(chan *pb.OrderCreatedEvent, 1)

	// simulate a subscriber (like merch service)
	bus.SubscribeToOrderCreated(func(event *pb.OrderCreatedEvent) {
		received <- event
	})

	orderItems := []*pb.OrderItem{
		{Sku: "food001", Price: 42, Qty: 1},
		{Sku: "drink001", Price: 2.5, Qty: 4},
	}

	// simulate publisher (like order service)
	testEvent := &pb.OrderCreatedEvent{
		EventId: "order-123",
		Order: &pb.Order{
			Id:         "order-123",
			UserId:     "user-abc",
			MerchantId: "merchant-xyz",
			Items:      orderItems,
			Status:     "created",
		},
	}

	if err := bus.PublishOrderCreated(testEvent); err != nil {
		t.Fatalf("Failed to publish event: %v", err)
	}

	receivedEvent := <-received

	if receivedEvent.Order.Id != "order-123" {
		t.Errorf("Expected order ID 'order-123', got '%s'", receivedEvent.Order.Id)
	}
	if receivedEvent.Order.Status != "created" {
		t.Errorf("Expected status 'created', got '%s'", receivedEvent.Order.Status)
	}
	if len(receivedEvent.Order.Items) != 2 || receivedEvent.Order.Items[1].Sku != "drink001" {
		t.Errorf("Unexpected items: %v", receivedEvent.Order.Items)
	}
}
