package category

import "time"

type Category struct {
	ID   int
	Name string
	// image string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "category"
}
