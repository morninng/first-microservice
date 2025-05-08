// order/internal/domain/repository/api.go
package repository

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/domain/model"
)

// OrderRepository defines the interface for order-related operations
type OrderRepository interface {
	// PlaceOrder creates a new order
	PlaceOrder(ctx context.Context, order model.Order) (model.Order, error)
	// GetOrder retrieves an order by its ID
	GetOrder(ctx context.Context, id int64) (model.Order, error)
}
