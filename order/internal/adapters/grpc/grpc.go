package grpc

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/application/core/domain"
	"github.com/morninng/first-proto/golang/order"
	log "github.com/sirupsen/logrus"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	log.WithContext(ctx).Info("Creating order...")
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	result, err := a.api.GetOrder(ctx, request.OrderId)
	var orderItems []*order.OrderItem
	for _, orderItem := range result.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	if err != nil {
		return nil, err
	}
	return &order.GetOrderResponse{UserId: result.CustomerID, OrderItems: orderItems}, nil
}
