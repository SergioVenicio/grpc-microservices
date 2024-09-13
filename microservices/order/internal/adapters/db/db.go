package db

import (
	"context"
	"fmt"

	"github.com/SergioVenicio/microservices/order/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float64
	Quantity    int64
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(ctx context.Context, id uint64) (domain.Order, error) {
	var orderEntity Order
	err := a.db.WithContext(ctx).First(&orderEntity, id).Error
	if err != nil {
		return domain.Order{}, err
	}

	var orderItems []domain.OrderItem
	err = a.db.WithContext(ctx).Find(&orderItems, "order_id=?", id).Error
	if err != nil {
		return domain.Order{}, err
	}
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.Unix(),
	}
	return order, nil
}

func (a Adapter) Save(ctx context.Context, order *domain.Order) error {
	var orderItems []OrderItem
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	err := a.db.WithContext(ctx).Create(&orderModel).Error
	if err != nil {
		return err
	}

	order.ID = int64(orderModel.ID)
	return nil
}

func NewAdapter(dbURI string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error %v", err)
	}

	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("payment"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error %v", err)
	}

	return &Adapter{db: db}, nil
}
