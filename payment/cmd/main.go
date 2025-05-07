package main

import (
	"log"

	"github.com/morninng/first-microservice/payment/config"
	"github.com/morninng/first-microservice/payment/internal/adapters/db"
	"github.com/morninng/first-microservice/payment/internal/adapters/grpc"
	"github.com/morninng/first-microservice/payment/internal/application/core/api"
)

const (
	service     = "payment"
	environment = "dev"
	id          = 2
)

func init() {
}

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
