package orderclient

import (
	"context"

	"github.com/L30Y3/nandemo/shared/models"
)

type OrderClientInterface interface {
	CreateOrder(ctx context.Context, order *models.Order) error
}
