package ports

import "github.com/SergioVenicio/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
