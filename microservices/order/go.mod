module github.com/SergioVenicio/microservices/order

go 1.22.2

require (
	github.com/huseyinbabal/microservices-proto/golang/order v1.0.39
	github.com/huseyinbabal/microservices-proto/golang/payment v1.0.39
	github.com/sirupsen/logrus v1.4.2
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.3.1
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.55.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.30.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.30.0
	go.opentelemetry.io/otel/sdk v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1
	google.golang.org/grpc v1.66.1
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.3.1 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
