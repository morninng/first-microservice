// order/internal/infrastructure/grpc/client/payment_client.go
package client

import (
	"context"

	"github.com/morninng/first-microservice/order/internal/domain/model"
	"github.com/morninng/first-proto/golang/payment"
)

// PaymentClient implements repository.PaymentRepository
type PaymentClient struct {
	client payment.PaymentClient
	conn   *Connection
}

// NewPaymentClient creates a new instance of PaymentClient
func NewPaymentClient(conn *Connection) *PaymentClient {
	return &PaymentClient{
		client: payment.NewPaymentClient(conn),
		conn:   conn,
	}
}

// Close closes the underlying gRPC connection
func (c *PaymentClient) Close() error {
	return c.conn.Close()
}

// Charge implements repository.PaymentRepository
func (c *PaymentClient) Charge(ctx context.Context, order *model.Order) error {
	_, err := c.client.Create(ctx, &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}
