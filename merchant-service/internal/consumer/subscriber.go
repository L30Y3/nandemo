package consumer

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/L30Y3/nandemo/shared/events"
	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

const (
	defaultProjectId = "nandemo-464411"
)

var firestoreClient *firestore.Client

func initFirestore(ctx context.Context, projectId string) {
	var err error
	firestoreClient, err = firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Yabai! Failed to initialise Firestore client: %v", err)
	}
}

// ListenForOrders simulates merchant listening for new orders
func ListenForOrders(bus events.EventBus) {
	ctx := context.Background()
	initFirestore(ctx, defaultProjectId)

	bus.SubscribeToOrderCreated(func(event *pb.OrderCreatedEvent) {
		log.Printf("Gambarimasu!! Handling OrderCreatedEvent: %+v", event)

		doc := map[string]interface{}{
			"orderId":     event.Order.Id,
			"userId":      event.Order.UserId,
			"merchantId":  event.Order.MerchantId, // currently there is only one merchant
			"items":       flattenOrderItems(event.Order.Items),
			"status":      event.Order.Status,
			"totalAmount": event.Order.TotalAmount,
			"createdAt":   event.Order.CreatedAt,
			"source":      event.Order.Source,
			"eventId":     event.EventId,
		}

		_, err := firestoreClient.Collection("orders").Doc(event.Order.Id).Set(ctx, doc)

		if err != nil {
			log.Printf("Yabai!! Failed to write order to Firestore: %v", err)
		} else {
			log.Printf("Yatta!! Order %s written to Firestore.", event.Order.Id)
		}
	})
}

func flattenOrderItems(items []*pb.OrderItem) []map[string]interface{} {
	flattened := make([]map[string]interface{}, len(items))
	for i, item := range items {
		flattened[i] = map[string]interface{}{
			"sku":   item.Sku,
			"qty":   item.Qty,
			"price": item.Price,
		}
	}
	return flattened
}
