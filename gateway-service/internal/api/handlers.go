package api

import (
	"encoding/json"
	"net/http"

	"github.com/L30Y3/nandemo/gateway-service/internal/oauth"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/login/google", oauth.GoogleLoginHandler)
	mux.HandleFunc("/login/microsoft", oauth.MicrosoftLoginHandler)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "Gateway Service OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
