package order

import (
	"crypto/rand"

	midtrans "github.com/veritrans/go-midtrans"
)

type Service interface {
	AddOrderDetail(input OrderDetailsInput) (OrderDetails, error)
	AddOrder(input OrderInput) (Order, error)
	GetOrderByID(ID int) (Order, error)
	GetPaymentURL(amount int64, name string, email string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddOrderDetail(input OrderDetailsInput) (OrderDetails, error) {
	orderDetails := OrderDetails{}
	orderDetails.OrderId = input.OrderId
	orderDetails.ProductName = input.ProductName
	orderDetails.ProductId = input.ProductId
	orderDetails.ProductQty = input.ProductQty
	orderDetails.ProductPrice = input.ProductPrice
	orderDetails.TotalProductPrice = input.TotalProductPrice

	newOrderDetails, err := s.repository.SaveOrderDetails(orderDetails)

	if err != nil {
		return newOrderDetails, err
	}

	return newOrderDetails, nil
}

func (s *service) AddOrder(input OrderInput) (Order, error) {
	order := Order{}
	order.UserId = input.UserId
	order.OrderTotalPrice = input.OrderTotalPrice
	order.OrderStatus = input.OrderStatus
	order.IsDelivered = input.IsDelivered
	order.PaymentURL = input.PaymentURL

	newOrder, err := s.repository.SaveOrder(order)

	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (s *service) GetOrderByID(ID int) (Order, error) {
	order, err := s.repository.FindByID(ID)

	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *service) GetPaymentURL(amount int64, name string, email string) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-X2E2PrnzkOHT6hqvOoJT0R1j"
	midclient.ClientKey = "SB-Mid-client-lLWn3SFH0KreU4Vy"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	randomBytes := make([]byte, 5)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  string(randomBytes),
			GrossAmt: amount,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: name,
			Email: email,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
