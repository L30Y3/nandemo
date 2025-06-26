package consumer

import (
	"fmt"

	"github.com/L30Y3/nandemo/shared/events"
	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

// ListenForOrders simulates merchant listening for new orders
func ListenForOrders(bus events.EventBus) {
	bus.SubscribeToOrderCreated(func(event *pb.OrderCreatedEvent) {
		fmt.Println("Merchant received new order:")
		fmt.Printf("-> Order ID: %s, User: %s, Items: %v\n",
			event.Order.Id, event.Order.UserId, event.Order.Items)
	})
}
