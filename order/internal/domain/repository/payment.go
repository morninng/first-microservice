// order/internal/domain/repository/payment.go
package repository

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/domain/model"
)

// PaymentRepository defines the interface for payment-related operations
type PaymentRepository interface {
	// Charge processes a payment for the given order
	Charge(ctx context.Context, order *model.Order) error
}
