package main

import (
	"log"

	"github.com/morninng/first-microservice/order/config"
	"github.com/morninng/first-microservice/order/internal/adapters/db"
	"github.com/morninng/first-microservice/order/internal/adapters/grpc"
	"github.com/morninng/first-microservice/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
