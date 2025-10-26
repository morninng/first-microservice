// order/internal/application/service/order_service.go
package service

import (
	"context"
	"strings"

	"github.com/morninng/first-microservice/order/internal/domain/model"
	"github.com/morninng/first-microservice/order/internal/domain/repository"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OrderService handles order-related business logic
type OrderService struct {
	db      repository.OrderPersistence
	payment repository.PaymentRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(db repository.OrderPersistence, payment repository.PaymentRepository) *OrderService {
	return &OrderService{
		db:      db,
		payment: payment,
	}
}

// PlaceOrder creates a new order and processes the payment
func (s *OrderService) PlaceOrder(ctx context.Context, order model.Order) (model.Order, error) {
	// Save the order to the database
	if err := s.db.Save(ctx, &order); err != nil {
		return model.Order{}, err
	}

	// Process the payment
	if err := s.payment.Charge(ctx, &order); err != nil {
		return s.handlePaymentError(err)
	}

	return order, nil
}

// GetOrder retrieves an order by its ID
func (s *OrderService) GetOrder(ctx context.Context, id int64) (model.Order, error) {
	return s.db.Get(ctx, id)
}

// handlePaymentError processes payment-related errors and converts them to appropriate gRPC status
func (s *OrderService) handlePaymentError(err error) (model.Order, error) {
	st := status.Convert(err)
	var allErrors []string

	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				allErrors = append(allErrors, violation.Description)
			}
		}
	}

	fieldErr := &errdetails.BadRequest_FieldViolation{
		Field:       "payment",
		Description: strings.Join(allErrors, "\n"),
	}

	badReq := &errdetails.BadRequest{}
	badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)

	orderStatus := status.New(codes.InvalidArgument, "order creation failed")
	statusWithDetails, _ := orderStatus.WithDetails(badReq)

	return model.Order{}, statusWithDetails.Err()
}
