// order/internal/infrastructure/persistence/order_repository.go
package persistence

import (
	"context"
	"fmt"

	"github.com/morninng/first-microservice/order/internal/domain/model"
	"gorm.io/gorm"
)

// OrderEntity represents the database model for an order
type OrderEntity struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItemEntity
}

// OrderItemEntity represents the database model for an order item
type OrderItemEntity struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

// OrderRepository implements repository.OrderPersistence
type OrderRepository struct {
	db *DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Get retrieves an order by its ID
func (r *OrderRepository) Get(ctx context.Context, id int64) (model.Order, error) {
	var orderEntity OrderEntity
	if err := r.db.WithContext(ctx).Preload("OrderItems").First(&orderEntity, id).Error; err != nil {
		return model.Order{}, fmt.Errorf("failed to get order: %w", err)
	}

	return r.toDomainModel(orderEntity), nil
}

// Save persists an order to the database
func (r *OrderRepository) Save(ctx context.Context, order *model.Order) error {
	orderEntity := r.toEntityModel(order)

	if err := r.db.WithContext(ctx).Create(&orderEntity).Error; err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}

	order.ID = int64(orderEntity.ID)
	return nil
}

// toDomainModel converts a database entity to a domain model
func (r *OrderRepository) toDomainModel(entity OrderEntity) model.Order {
	orderItems := make([]model.OrderItem, len(entity.OrderItems))
	for i, item := range entity.OrderItems {
		orderItems[i] = model.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		}
	}

	return model.Order{
		ID:         int64(entity.ID),
		CustomerID: entity.CustomerID,
		Status:     model.OrderStatus(entity.Status),
		OrderItems: orderItems,
		CreatedAt:  entity.CreatedAt.Unix(),
	}
}

// toEntityModel converts a domain model to a database entity
func (r *OrderRepository) toEntityModel(order *model.Order) OrderEntity {
	orderItems := make([]OrderItemEntity, len(order.OrderItems))
	for i, item := range order.OrderItems {
		orderItems[i] = OrderItemEntity{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		}
	}

	return OrderEntity{
		CustomerID: order.CustomerID,
		Status:     string(order.Status),
		OrderItems: orderItems,
	}
}
