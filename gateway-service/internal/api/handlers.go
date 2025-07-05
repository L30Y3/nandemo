package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/L30Y3/nandemo/gateway-service/internal/oauth"
	orderclient "github.com/L30Y3/nandemo/shared/clients/orderclient"
)

const (
	healthRoute            = "/health"
	loginGoogleRoute       = "/login/google"
	loginMicrosoftRoute    = "/login/microsoft"
	orderRoute             = "/order"
	getMerchantOrdersRoute = "/merchant/{merchantId}/orders"
	getMerchantGoodsRoute  = "/merchant/{merchantId}/goods"
)

type HealthResponse struct {
	Status string `json:"status"`
}

var orderSvc = orderclient.NewOrderServiceClient()

func RegisterRoutes(r chi.Router) {
	r.Get(healthRoute, HealthHandler)
	r.Get(loginGoogleRoute, oauth.GoogleLoginHandler)
	r.Get(loginMicrosoftRoute, oauth.MicrosoftLoginHandler)
	r.Post(orderRoute, HandleCreateOrder)
	r.Get(getMerchantGoodsRoute, HandleGetMerchantGoods)
	r.Get(getMerchantOrdersRoute, HandleGetMerchantOrders)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "Gateway Service OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// TODO: forward to order-service
	w.Write([]byte("Order created (stub)"))
}

func HandleGetMerchantOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: forward to merchant-service
	w.Write([]byte("Merchant orders (stub)"))
}

func HandleGetMerchantGoods(w http.ResponseWriter, r *http.Request) {
	// TODO: forward to merchant-service
	w.Write([]byte("Merchant goods (stub)"))
}
