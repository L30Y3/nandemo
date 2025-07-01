package models

type Order struct {
	ID         string      `json:"id"`
	UserID     string      `json:"user_id"`
	MerchantID string      `json:"merchant_id"`
	Items      []OrderItem `json:"items"`  // Simplified: product/service item names
	Status     string      `json:"status"` // e.g. "created", "confirmed", "fulfilled"
}

type OrderItem struct {
	SKU   string  `json:"sku"`
	Qty   int32   `json:"qty"`
	Price float64 `json:"price"`
}
