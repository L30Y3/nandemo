package merchantclient

import (
	"context"

	"github.com/L30Y3/nandemo/shared/models"
)

type MerchantClientInterface interface {
	GetMerchantGoods(ctx context.Context, merchantID string) ([]models.Goods, error)
	GetMerchantOrdersWithWindow(ctx context.Context, merchantID string, window string) ([]models.Order, error)
}
