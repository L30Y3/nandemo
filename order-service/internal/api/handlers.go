package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/L30Y3/nandemo/shared/events"
	"github.com/L30Y3/nandemo/shared/models"
	pb "github.com/L30Y3/nandemo/shared/proto/protoevents"
)

var eventBus events.EventBus

func RegisterRoutes(mux *http.ServeMux, bus events.EventBus) {
	eventBus = bus
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/order", CreateOrderHandler)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"status": "Order Service OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// For now: log order, simulate successful processing
	fmt.Printf("Received new order: %+v\n", order)

	// create Protobuf OrderCreatedEvent from incoming order model
	event := &pb.OrderCreatedEvent{
		EventId: "evt-" + order.ID,
		Order: &pb.Order{
			Id:         order.ID,
			UserId:     order.UserID,
			MerchantId: order.MerchantID,
			Items:      order.Items,
			Status:     order.Status,
		},
	}

	if err := eventBus.PublishOrderCreated(event); err != nil {
		fmt.Printf("Failed to publish eevnt: %v\n", err)
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
