package product

type ProductFormatter struct {
	ID          int     `json:"id"`
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:          product.ID,
		CategoryId:  product.CategoryId,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	return formatter
}

type ListProductFormatter struct {
	ID         int     `json:"id"`
	CategoryId int     `json:"category_id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
}

func FormatListProduct(product Product) ListProductFormatter {
	formatter := ListProductFormatter{
		ID:         product.ID,
		CategoryId: product.CategoryId,
		Name:       product.Name,
		Price:      product.Price,
	}

	return formatter
}

func ListFormatProducts(product []Product) []ListProductFormatter {
	if len(product) == 0 {
		return []ListProductFormatter{}
	}

	var listProductFormatter []ListProductFormatter

	for _, product := range product {
		formatter := FormatListProduct(product)
		listProductFormatter = append(listProductFormatter, formatter)
	}

	return listProductFormatter
}
