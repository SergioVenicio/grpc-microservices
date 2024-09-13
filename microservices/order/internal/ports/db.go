package ports

import (
	"context"

	"github.com/SergioVenicio/microservices/order/internal/application/core/domain"
)

type DBPort interface {
	Get(context.Context, uint64) (domain.Order, error)
	Save(context.Context, *domain.Order) error
}
