package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/L30Y3/nandemo/shared/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type mockMerchantClient struct{}
type mockOrderClient struct{}

func (m *mockMerchantClient) GetMerchantGoods(ctx context.Context, merchantID string) ([]models.Goods, error) {
	return []models.Goods{
		{SKU: "123", Name: "Test Goods", Price: 10.99, Category: "Test Category", Description: "Test Description", Available: true}}, nil
}

func (m *mockMerchantClient) GetMerchantOrdersWithWindow(ctx context.Context, merchantID string, window string) ([]models.Order, error) {
	// Return a mock order for testing
	if window == "12h" {
		return []models.Order{
			{ID: "1",
				MerchantID: merchantID,
				UserID:     "user1",
				Items:      []models.OrderItem{{SKU: "123", Qty: 2, Price: 10.99}},
				Status:     "pending"},
		}, nil
	} else {
		return []models.Order{}, nil
	}
}

func (m *mockOrderClient) CreateOrder(ctx context.Context, order *models.Order) error {
	return nil
}

func TestHandleGetMerchantGoods(t *testing.T) {
	h := &Handler{
		MerchantClient: &mockMerchantClient{},
		OrderClient:    &mockOrderClient{},
	}

	// Create a mock HTTP request
	req := httptest.NewRequest("GET", "/merchant/123/goods", nil)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/merchant/{merchantId}/goods", h.HandleGetMerchantGoods)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200 OK")

	var goods []models.Goods
	err := json.Unmarshal(rr.Body.Bytes(), &goods)
	assert.NoError(t, err, "Expected no error while unmarshalling response")
	assert.Len(t, goods, 1, "Expected one goods item in response")
	assert.Equal(t, "Test Goods", goods[0].Name, "Expected goods name to be 'Test Goods'")
}

func TestCreateOrder(t *testing.T) {
	h := &Handler{
		MerchantClient: &mockMerchantClient{},
		OrderClient:    &mockOrderClient{},
	}

	order := models.Order{
		ID:         "1",
		UserID:     "user1",
		MerchantID: "merchant1",
		Items:      []models.OrderItem{{SKU: "123", Qty: 2, Price: 10.99}},
		Status:     "created",
	}

	reqBody, _ := json.Marshal(order)
	req := httptest.NewRequest("POST", "/order", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Post("/order", h.HandleCreateOrder)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201 Created")

	var createdOrder models.Order
	err := json.Unmarshal(rr.Body.Bytes(), &createdOrder)
	assert.NoError(t, err, "Expected no error while unmarshalling response")
	assert.Equal(t, order.UserID, createdOrder.UserID, "Expected order user IDs to match")
}
