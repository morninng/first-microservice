// order/internal/domain/model/order_item.go
package model

import "errors"

// OrderItem represents an item in an order
type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

// NewOrderItem creates a new order item
func NewOrderItem(productCode string, unitPrice float32, quantity int32) (*OrderItem, error) {
	if productCode == "" {
		return nil, errors.New("product code is required")
	}
	if unitPrice <= 0 {
		return nil, errors.New("unit price must be greater than zero")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	return &OrderItem{
		ProductCode: productCode,
		UnitPrice:   unitPrice,
		Quantity:    quantity,
	}, nil
}

// TotalPrice calculates the total price of the order item
func (i *OrderItem) TotalPrice() float32 {
	return i.UnitPrice * float32(i.Quantity)
}

// Validate validates the order item
func (i *OrderItem) Validate() error {
	if i.ProductCode == "" {
		return errors.New("product code is required")
	}
	if i.UnitPrice <= 0 {
		return errors.New("unit price must be greater than zero")
	}
	if i.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	return nil
}
