// order/internal/adapters/grpc/server/server.go
package server

import (
	"fmt"
	"net"

	"github.com/morninng/first-microservice/order/config"
	"github.com/morninng/first-microservice/order/internal/domain/repository"
	"github.com/morninng/first-proto/golang/order"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// OrderServer handles gRPC requests for the order service
type OrderServer struct {
	api    repository.OrderRepository
	port   int
	server *grpc.Server
	order.UnimplementedOrderServer
}

// NewOrderServer creates a new instance of OrderServer
func NewOrderServer(api repository.OrderRepository, port int) *OrderServer {
	return &OrderServer{
		api:  api,
		port: port,
	}
}

// Start initializes and starts the gRPC server
func (s *OrderServer) Start() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", s.port, err)
	}

	s.server = grpc.NewServer()
	order.RegisterOrderServer(s.server, s)

	if config.GetEnv() == "development" {
		reflection.Register(s.server)
	}

	log.Printf("starting order service on port %d ...", s.port)
	if err := s.server.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve grpc: %w", err)
	}

	return nil
}

// Stop gracefully shuts down the gRPC server
func (s *OrderServer) Stop() {
	if s.server != nil {
		s.server.Stop()
	}
}
