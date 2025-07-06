package api

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/L30Y3/nandemo/shared/models"
	"github.com/go-chi/chi/v5"
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
}

func (h *MerchantHandlerWithFirestoreClient) GetMerchantGoodsHandler(w http.ResponseWriter, r *http.Request) {
	merchantId := chi.URLParam(r, "merchantId")
	ctx := r.Context()

	goods, err := GetGoodsByMerchant(ctx, h.Firestore, merchantId)
	if err != nil {
		http.Error(w, "Failed to fetch goods", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
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

		goodsList[i] = models.Goods{
			SKU:         gMap["sku"].(string),
			Name:        gMap["name"].(string),
			Price:       gMap["price"].(float64),
			Category:    "", // optional if not stored
			Description: "",
			Available:   true, // assuming always available for now
		}
	}

	return goodsList, nil
}
