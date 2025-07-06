// shared/models/goods.go

package models

type Goods struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Available   bool    `json:"available"`
}
