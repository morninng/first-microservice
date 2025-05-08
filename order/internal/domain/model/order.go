// order/internal/domain/model/order.go
package model

import (
	"errors"
	"time"
)

// OrderStatus represents the current status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

// Order represents an order in the system
type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     OrderStatus `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
	UpdatedAt  int64       `json:"updated_at"`
}

// NewOrder creates a new order with the given customer ID and items
func NewOrder(customerID int64, orderItems []OrderItem) (*Order, error) {
	if customerID <= 0 {
		return nil, errors.New("invalid customer ID")
	}
	if len(orderItems) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	now := time.Now().Unix()
	return &Order{
		CustomerID: customerID,
		Status:     OrderStatusPending,
		OrderItems: orderItems,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

// TotalPrice calculates the total price of the order
func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.TotalPrice()
	}
	return totalPrice
}

// MarkAsPaid marks the order as paid
func (o *Order) MarkAsPaid() error {
	if o.Status != OrderStatusPending {
		return errors.New("only pending orders can be marked as paid")
	}
	o.Status = OrderStatusPaid
	o.UpdatedAt = time.Now().Unix()
	return nil
}

// Cancel cancels the order
func (o *Order) Cancel() error {
	if o.Status != OrderStatusPending {
		return errors.New("only pending orders can be cancelled")
	}
	o.Status = OrderStatusCancelled
	o.UpdatedAt = time.Now().Unix()
	return nil
}

// Validate validates the order
func (o *Order) Validate() error {
	if o.CustomerID <= 0 {
		return errors.New("invalid customer ID")
	}
	if len(o.OrderItems) == 0 {
		return errors.New("order must have at least one item")
	}
	for _, item := range o.OrderItems {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
