package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/SergioVenicio/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices-proto/golang/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	var orderItems []domain.OrderItem

	for _, item := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   float64(item.UnitPrice),
			Quantity:    int64(item.Quantity),
		})
	}
	newOrder := domain.NewOder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		err = status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("failed to charge user: %d", request.UserId),
		)
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	var orderItems []*order.OrderItem

	dbOrder, err := a.api.GetOrder(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	for _, item := range dbOrder.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductCode: item.ProductCode,
			Quantity:    int32(item.Quantity),
			UnitPrice:   float32(item.UnitPrice),
		})
	}

	return &order.GetOrderResponse{
		UserId:     dbOrder.CustomerID,
		OrderItems: orderItems,
	}, nil
}
