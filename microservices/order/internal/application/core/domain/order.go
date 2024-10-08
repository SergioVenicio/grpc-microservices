package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int64   `json:"quantity"`
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

func (o Order) TotalPrice() float32 {
	var total float32

	for _, item := range o.OrderItems {
		total += float32(item.Quantity) * float32(item.UnitPrice)
	}
	return total
}

func NewOder(customerID int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerID,
		OrderItems: orderItems,
	}
}
