module github.com/morninng/first-microservice/payment

go 1.23.0

toolchain go1.23.9

require (
	github.com/morninng/first-proto/golang/payment v0.0.1
	github.com/sirupsen/logrus v1.9.3
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.3.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.60.0
	go.opentelemetry.io/otel v1.35.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.34.0
	go.opentelemetry.io/otel/trace v1.35.0
	google.golang.org/grpc v1.72.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.26.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.3.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250505200425-f936aa4a68b2 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
