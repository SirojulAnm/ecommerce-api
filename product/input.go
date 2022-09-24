package product

type ProductInput struct {
	CategoryId  int     `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

type IdInput struct {
	ID int `uri:"id" binding:"required"`
}
