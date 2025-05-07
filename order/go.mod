module github.com/morninng/first-microservice/order

go 1.23.0

toolchain go1.23.8

require (
	github.com/morninng/first-proto/golang/order v1.2.8
	github.com/morninng/first-proto/golang/payment v0.0.1
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34
	google.golang.org/grpc v1.72.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.26.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
