package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/L30Y3/nandemo/shared/models"
	"github.com/go-chi/chi/v5"
	"google.golang.org/api/iterator"
)

const (
	getGoodiesRoute = "/merchant/{merchantId}/goods"
	getOrdersRoute  = "/merchant/{merchantId}/orders"
)

type MerchantHandlerWithFirestoreClient struct {
	Firestore *firestore.Client
}

func RegisterRoutes(r chi.Router, h *MerchantHandlerWithFirestoreClient) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Merchant Service OK"))
	})
	r.Get(getGoodiesRoute, h.GetMerchantGoodsHandler)
	r.Get(getOrdersRoute, h.GetMerchantOrdersHandler)
}

func (h *MerchantHandlerWithFirestoreClient) GetMerchantGoodsHandler(w http.ResponseWriter, r *http.Request) {
	merchantId := chi.URLParam(r, "merchantId")
	ctx := r.Context()

	goods, err := GetGoodsByMerchant(ctx, h.Firestore, merchantId)
	if err != nil {
		log.Printf("Failed to fetch goods for merchantId=%s: %v", merchantId, err)
		http.Error(w, "Failed to fetch goods", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
}

func (h *MerchantHandlerWithFirestoreClient) GetMerchantOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get merchantId from the URL
	merchantID := chi.URLParam(r, "merchantId")

	// Get the `window` query param (e.g., "12h")
	windowParam := r.URL.Query().Get("window")
	if windowParam == "" {
		windowParam = "12h" // default
	}

	windowDuration, err := time.ParseDuration(windowParam)
	if err != nil {
		http.Error(w, "Invalid window format", http.StatusBadRequest)
		return
	}

	cutoff := time.Now().Add(-windowDuration)

	// Fetch from Firestore
	orders, err := GetOrdersSince(ctx, merchantID, cutoff, h.Firestore)
	if err != nil {
		log.Printf("Failed to fetch orders for merchantId=%s: %v", merchantID, err)
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetGoodsByMerchant(ctx context.Context, fs *firestore.Client, merchantId string) ([]models.Goods, error) {
	docRef := fs.Collection("merchants").Doc(merchantId)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	data := doc.Data()
	goodsRaw, ok := data["goodies"]
	if !ok {
		return nil, nil
	}

	// convert []interface{} to []models.Goods
	goodsSlice := goodsRaw.([]interface{})
	goodsList := make([]models.Goods, len(goodsSlice))

	for i, g := range goodsSlice {
		gMap := g.(map[string]interface{})
		log.Printf("Got good: %+v", gMap)
		var price float64
		switch v := gMap["price"].(type) {
		case int64:
			price = float64(v)
		case float64:
			price = v
		case string:
			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid string price: %v", v)
			}
			price = parsed
		default:
			return nil, fmt.Errorf("unexpected type for price: %T", v)
		}

		goodsList[i] = models.Goods{
			SKU:         gMap["sku"].(string),
			Name:        gMap["name"].(string),
			Price:       price,
			Category:    "", // optional if not stored
			Description: "",
			Available:   true, // assuming always available for now
		}
	}

	return goodsList, nil
}

func GetOrdersSince(ctx context.Context, merchantID string, cutoff time.Time, fs *firestore.Client) ([]models.Order, error) {
	iter := fs.Collection("orders").
		Where("merchantId", "==", merchantID).
		Where("createdAt", ">", cutoff).
		Documents(ctx)

	defer iter.Stop()

	var orders []models.Order

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("firestore iteration error: %w", err)
		}

		var order models.Order
		if err := doc.DataTo(&order); err != nil {
			return nil, fmt.Errorf("firestore decode error: %w", err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
