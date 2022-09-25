package order

type CartInput struct {
	ProductId  int `json:"product_id" binding:"required"`
	ProductQty int `json:"product_qty" binding:"required"`
}

type ShoppingCartInput struct {
	Cart []CartInput `form:"cart[]" binding:"dive"`
}

type OrderDetailsInput struct {
	OrderId           int
	ProductId         int
	ProductName       string
	ProductQty        int
	ProductPrice      float32
	TotalProductPrice float32
}

type OrderInput struct {
	UserId          int
	OrderTotalPrice float32
	OrderStatus     string
	IsDelivered     bool
	PaymentURL      string
}
