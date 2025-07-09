package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/L30Y3/nandemo/gateway-service/internal/oauth"
	merchantclient "github.com/L30Y3/nandemo/shared/clients/merchantclient"
	orderclient "github.com/L30Y3/nandemo/shared/clients/orderclient"
	"github.com/L30Y3/nandemo/shared/models"
)

const (
	healthRoute            = "/health"
	loginGoogleRoute       = "/login/google"
	loginMicrosoftRoute    = "/login/microsoft"
	orderRoute             = "/order"
	getMerchantOrdersRoute = "/merchant/{merchantId}/orders"
	getMerchantGoodsRoute  = "/merchant/{merchantId}/goods"
)

type Handler struct {
	OrderClient    orderclient.OrderClientInterface
	MerchantClient merchantclient.MerchantClientInterface
}

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get(healthRoute, h.HealthHandler)
	r.Get(loginGoogleRoute, oauth.GoogleLoginHandler)
	r.Get(loginMicrosoftRoute, oauth.MicrosoftLoginHandler)
	r.Post(orderRoute, h.HandleCreateOrder)
	r.Get(getMerchantGoodsRoute, h.HandleGetMerchantGoods)
	r.Get(getMerchantOrdersRoute, h.HandleGetMerchantOrders)
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "Gateway Service OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.OrderClient.CreateOrder(r.Context(), &order); err != nil {
		http.Error(w, "Failed to forward order", http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *Handler) HandleGetMerchantOrders(w http.ResponseWriter, r *http.Request) {
	merchantId := chi.URLParam(r, "merchantId")
	window := r.URL.Query().Get("window")

	if window == "" {
		window = "12h"
	}

	orders, err := h.MerchantClient.GetMerchantOrdersWithWindow(r.Context(), merchantId, window)
	if err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) HandleGetMerchantGoods(w http.ResponseWriter, r *http.Request) {
	merchantId := chi.URLParam(r, "merchantId")

	goods, err := h.MerchantClient.GetMerchantGoods(r.Context(), merchantId)
	if err != nil {
		http.Error(w, "Failed to fetch goods", http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
}
