package main

import (
	"log"

	"github.com/morninng/first-microservice/order/config"
	"github.com/morninng/first-microservice/order/internal/adapters/grpc/server"
	"github.com/morninng/first-microservice/order/internal/application/service"
	"github.com/morninng/first-microservice/order/internal/infrastructure/grpc/client"
	"github.com/morninng/first-microservice/order/internal/infrastructure/persistence"
)

func main() {

	// データベース接続の初期化
	db, err := persistence.NewDB(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	// リポジトリの初期化
	orderRepository := persistence.NewOrderRepository(db)

	// gRPCクライアントの初期化
	paymentConn, err := client.NewConnection(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}
	defer paymentConn.Close()
	paymentClient := client.NewPaymentClient(paymentConn)

	application := service.NewOrderService(orderRepository, paymentClient)
	grpcAdapter := server.NewOrderServer(application, config.GetApplicationPort())
	grpcAdapter.Start()
}
