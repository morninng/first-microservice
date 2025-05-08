package ports

import (
	"context"

	"github.com/morninng/first-microservice/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
