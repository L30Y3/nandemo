package orderclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/L30Y3/nandemo/shared/config"
	"github.com/L30Y3/nandemo/shared/models"
)

type OrderServiceClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewOrderServiceClient() *OrderServiceClient {
	return &OrderServiceClient{
		BaseURL:    getBaseURL(),
		HTTPClient: http.DefaultClient,
	}
}

func (c *OrderServiceClient) CreateOrder(ctx context.Context, order *models.Order) error {
	prefix := "[order client]:"
	url := fmt.Sprintf("%s/order", c.BaseURL)

	body, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("%s failed to marshal order: %w", prefix, err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("%s failed to build request: %w", prefix, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s request failed: %w", prefix, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("%s order service returned status: %s", prefix, resp.Status)
	}

	return nil
}

func getBaseURL() string {
	return fmt.Sprintf("http://localhost:%s", config.OrderServicePort)
}
