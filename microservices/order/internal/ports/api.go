package ports

import (
	"context"

	"github.com/SergioVenicio/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	GetOrder(context.Context, int64) (*domain.Order, error)
	PlaceOrder(context.Context, domain.Order) (domain.Order, error)
}
