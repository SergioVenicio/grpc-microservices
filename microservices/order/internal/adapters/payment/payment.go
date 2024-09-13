package payment

import (
	"context"
	"time"

	"github.com/SergioVenicio/microservices/order/internal/application/core/domain"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/huseyinbabal/microservices-proto/golang/payment"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}

func NewAdapter(paymentServiceURL string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponentialWithJitter(time.Second, 0.10)),
	)))
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	conn, err := grpc.NewClient(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, err
}
