package models

type Order struct {
	ID         string   `json:"id"`
	UserID     string   `json:"user_id"`
	MerchantID string   `json:"merchant_id"`
	Items      []string `json:"items"`  // Simplified: product/service item names
	Status     string   `json:"status"` // e.g. "created", "confirmed", "fulfilled"
}
