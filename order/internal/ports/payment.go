package ports

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
