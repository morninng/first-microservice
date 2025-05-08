package client

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Connection represents a gRPC connection
type Connection struct {
	*grpc.ClientConn
}

// NewConnection creates a new gRPC connection
func NewConnection(address string) (*Connection, error) {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return &Connection{conn}, nil
}

// Close closes the gRPC connection
func (c *Connection) Close() error {
	return c.ClientConn.Close()
}
