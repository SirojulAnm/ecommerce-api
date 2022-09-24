package product

import "time"

type Product struct {
	ID          int
	CategoryId  int
	Name        string
	Description string
	Price       float32
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tabler interface {
	TableName() string
}

func (Product) TableName() string {
	return "product"
}
