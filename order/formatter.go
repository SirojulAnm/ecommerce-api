package order

type OrderFormatter struct {
	ID              int                     `json:"id"`
	UserId          int                     `json:"user_id"`
	OrderTotalPrice float32                 `json:"order_total_price"`
	OrderStatus     string                  `json:"order_status"`
	IsDelivered     bool                    `json:"is_delivered"`
	PaymentURL      string                  `json:"payment_url"`
	OrderDetails    []OrderDetailsFormatter `json:"order_details"`
}

type OrderDetailsFormatter struct {
	ID                int     `json:"id"`
	OrderId           int     `json:"order_id"`
	ProductId         int     `json:"product_id"`
	ProductName       string  `json:"product_name"`
	ProductQty        int     `json:"product_qty"`
	ProductPrice      float32 `json:"product_price"`
	TotalProductPrice float32 `json:"total_product_price"`
}

func FormatOrder(order Order) OrderFormatter {
	campaignDetailFormatter := OrderFormatter{}
	campaignDetailFormatter.ID = order.ID
	campaignDetailFormatter.UserId = order.UserId
	campaignDetailFormatter.OrderTotalPrice = order.OrderTotalPrice
	campaignDetailFormatter.OrderStatus = order.OrderStatus
	campaignDetailFormatter.IsDelivered = order.IsDelivered
	campaignDetailFormatter.PaymentURL = order.PaymentURL

	var orderDetailsFormatter []OrderDetailsFormatter
	if len(order.OrderDetails) > 0 {
		for _, orderItem := range order.OrderDetails {
			formatter := FormatOrderDetails(orderItem)
			orderDetailsFormatter = append(orderDetailsFormatter, formatter)
		}
	}

	campaignDetailFormatter.OrderDetails = orderDetailsFormatter

	return campaignDetailFormatter
}

func FormatOrderDetails(orderDetails OrderDetails) OrderDetailsFormatter {
	orderDetailsFormatter := OrderDetailsFormatter{}
	orderDetailsFormatter.ID = orderDetails.ID
	orderDetailsFormatter.OrderId = orderDetails.OrderId
	orderDetailsFormatter.ProductId = orderDetails.ProductId
	orderDetailsFormatter.ProductName = orderDetails.ProductName
	orderDetailsFormatter.ProductQty = orderDetails.ProductQty
	orderDetailsFormatter.ProductPrice = orderDetails.ProductPrice
	orderDetailsFormatter.TotalProductPrice = orderDetails.TotalProductPrice

	return orderDetailsFormatter
}
