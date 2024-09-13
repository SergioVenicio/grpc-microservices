package main

import (
	"context"
	"os"
	"time"

	"github.com/SergioVenicio/microservices/order/config"
	"github.com/SergioVenicio/microservices/order/internal/adapters/db"
	"github.com/SergioVenicio/microservices/order/internal/adapters/grpc"
	"github.com/SergioVenicio/microservices/order/internal/adapters/payment"
	"github.com/SergioVenicio/microservices/order/internal/application/core/api"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	service     = "order"
	environment = "dev"
	id          = 1
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(url),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, err
	}
	tracerprovider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(
			exporter,
			tracesdk.WithMaxExportBatchSize(tracesdk.DefaultMaxExportBatchSize),
			tracesdk.WithBatchTimeout(tracesdk.DefaultScheduleDelay*time.Millisecond),
			tracesdk.WithMaxExportBatchSize(tracesdk.DefaultMaxExportBatchSize),
		),
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(service),
				attribute.String("environment", environment),
				attribute.Int64("ID", id),
			),
		),
	)
	return tracerprovider, nil
}

func main() {
	tp, err := tracerProvider("localhost:4318")
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database. Error: %v", err)
	}

	payment, err := payment.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("failed to initialize payment stub. Error %v", err)
	}

	app := api.NewApplication(dbAdapter, payment)
	grpc := grpc.NewAdapter(app, config.GetApplicationPort())
	grpc.Run()
}

type customLogger struct {
	formatter log.JSONFormatter
}

func (l customLogger) Format(entry *log.Entry) ([]byte, error) {
	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	//Below injection is Just to understand what Context has
	entry.Data["Context"] = span.SpanContext()
	return l.formatter.Format(entry)
}

func init() {
	log.SetFormatter(customLogger{
		formatter: log.JSONFormatter{FieldMap: log.FieldMap{
			"msg": "message",
		}},
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
