// order/internal/domain/repository/db.go
package repository

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/domain/model"
)

// OrderPersistence defines the interface for order persistence operations
type OrderPersistence interface {
	// Save persists an order to the database
	Save(ctx context.Context, order *model.Order) error
	// Get retrieves an order from the database by its ID
	Get(ctx context.Context, id int64) (model.Order, error)
}
