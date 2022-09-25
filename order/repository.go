package order

import "gorm.io/gorm"

type Repository interface {
	SaveOrderDetails(orderDetails OrderDetails) (OrderDetails, error)
	SaveOrder(order Order) (Order, error)
	FindByID(ID int) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveOrderDetails(orderDetails OrderDetails) (OrderDetails, error) {
	err := r.db.Create(&orderDetails).Error
	if err != nil {
		return orderDetails, err
	}

	return orderDetails, nil
}

func (r *repository) SaveOrder(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) FindByID(ID int) (Order, error) {
	var order Order
	err := r.db.Preload("OrderDetails").Find(&order, ID).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
