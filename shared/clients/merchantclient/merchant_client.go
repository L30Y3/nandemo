package merchantclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/L30Y3/nandemo/shared/config"
	"github.com/L30Y3/nandemo/shared/models"
)

type MerchantServiceClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func getBaseURL() string {
	return fmt.Sprintf("http://localhost:%s", config.MerchantServicePort)
}

func NewMerchantServiceClient() *MerchantServiceClient {
	return &MerchantServiceClient{
		BaseURL:    getBaseURL(),
		HTTPClient: http.DefaultClient,
	}
}

func (c *MerchantServiceClient) GetMerchantGoods(ctx context.Context, merchantId string) ([]models.Goods, error) {
	prefix := "[merchant client]:"
	url := fmt.Sprintf("%s/merchant/%s/goods", c.BaseURL, merchantId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s failed to build request: %w", prefix, err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s HTTP request failed: %w", prefix, err)
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("%s merchant service returned empty response body", prefix)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s merchant service returned status %d: %s", prefix, resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Printf("%s raw body response: %s", prefix, string(bodyBytes))

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var goods []models.Goods
	if err := json.NewDecoder(resp.Body).Decode(&goods); err != nil {
		return nil, fmt.Errorf("%s failed to decode goods response: %w", prefix, err)
	}

	return goods, nil
}

func (c *MerchantServiceClient) GetMerchantOrdersWithWindow(ctx context.Context, merchantId, window string) ([]models.Order, error) {
	prefix := "[merchant client]:"
	url := fmt.Sprintf("%s/merchant/%s/orders?window=%s", c.BaseURL, merchantId, window)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s failed to build request: %w", prefix, err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s request failed: %w", prefix, err)
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("%s merchant service returned empty response body", prefix)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(req.Body)
		return nil, fmt.Errorf("%s merchant service returned status %d: %s", prefix, resp.StatusCode, string(bodyBytes))
	}

	var orders []models.Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, fmt.Errorf("failed to decode orders: %w", err)
	}

	return orders, nil
}
