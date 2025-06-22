package events

import "github.com/L30Y3/nandemo/shared/models"

type OrderCreatedEvent struct {
	EventID string       `json:"event_id"`
	Order   models.Order `json:"order"`
}
