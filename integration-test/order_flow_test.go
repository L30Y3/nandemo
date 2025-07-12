package integrationtest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/L30Y3/nandemo/shared/models"
)

func TestOrderLifecycle(t *testing.T) {
	order := models.Order{
		ID:         "test-order-001",
		UserID:     "u01",
		MerchantID: "m01",
		Items: []models.OrderItem{
			{SKU: "123", Qty: 1, Price: 10.0},
		},
		Status: "created",
	}

	// Step 1: POST /order
	body, err := json.Marshal(order)
	require.NoError(t, err)

	resp, err := http.Post("http://localhost:8080/order", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Step 2: Poll /merchant/{id}/orders?window=12h
	var fetched []models.Order
	maxRetries := 10
	interval := time.Second

	for i := 0; i < maxRetries; i++ {
		res, err := http.Get("http://localhost:8080/merchant/m01/orders?window=12h")
		require.NoError(t, err)

		bodyBytes, _ := io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Logf("[Retry %d/%d] Unexpected status: %d: %s", i+1, maxRetries, res.StatusCode, string(bodyBytes))
			time.Sleep(interval)
			continue
		}

		err = json.Unmarshal(bodyBytes, &fetched)
		require.NoError(t, err)

		if len(fetched) > 0 {
			break
		}
		t.Logf("[retry %d/%d] No orders yet, waiting...", i+1, maxRetries)
		time.Sleep(interval)
	}

	require.NotEmpty(t, fetched)
	assert.Equal(t, order.ID, fetched[0].ID)
	fmt.Println("Order flow integration test passed.")
}
