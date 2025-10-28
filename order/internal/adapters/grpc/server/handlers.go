// order/internal/adapters/grpc/server/handlers.go
package server

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/domain/model"
	"github.com/morninng/first-proto/golang/order"
	log "github.com/sirupsen/logrus"
)

// Create handles the creation of a new order
func (s *OrderServer) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	log.WithContext(ctx).Info("Creating order... aaa bbb ")

	orderItems := convertToOrderItems(request.OrderItems)
	newOrder, err := model.NewOrder(request.UserId, orderItems)
	if err != nil {
		return nil, err
	}

	result, err := s.api.PlaceOrder(ctx, *newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		OrderId: result.ID,
	}, nil
}

// Get retrieves an order by its ID
func (s *OrderServer) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	result, err := s.api.GetOrder(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return &order.GetOrderResponse{
		UserId:     result.CustomerID,
		OrderItems: convertToProtoOrderItems(result.OrderItems),
	}, nil
}

// convertToOrderItems converts proto order items to domain order items
func convertToOrderItems(items []*order.OrderItem) []model.OrderItem {
	orderItems := make([]model.OrderItem, len(items))
	for i, item := range items {
		orderItems[i] = model.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		}
	}
	return orderItems
}

// convertToProtoOrderItems converts domain order items to proto order items
func convertToProtoOrderItems(items []model.OrderItem) []*order.OrderItem {
	protoItems := make([]*order.OrderItem, len(items))
	for i, item := range items {
		protoItems[i] = &order.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		}
	}
	return protoItems
}
