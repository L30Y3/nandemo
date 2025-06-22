package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/L30Y3/nandemo/shared/models"
)

func RegisterRoutes(mux *http.ServeMux) {
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
