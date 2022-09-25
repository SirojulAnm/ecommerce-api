package order

import "time"

type OrderDetails struct {
	ID                int
	OrderId           int
	ProductId         int
	ProductName       string
	ProductQty        int
	ProductPrice      float32
	TotalProductPrice float32
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Order struct {
	ID              int
	UserId          int
	OrderTotalPrice float32
	OrderStatus     string
	IsDelivered     bool
	PaymentURL      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	OrderDetails    []OrderDetails
}

type Tabler interface {
	TableName() string
}

func (OrderDetails) TableName() string {
	return "order_details"
}

func (Order) TableName() string {
	return "order"
}
