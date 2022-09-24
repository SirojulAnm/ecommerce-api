package product

type Service interface {
	AddProduct(input ProductInput) (Product, error)
	GetProductDetailByID(ID int) (Product, error)
	ListProductsByCategoryID(categoryID int) ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddProduct(input ProductInput) (Product, error) {
	product := Product{}
	product.CategoryId = input.CategoryId
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) GetProductDetailByID(ID int) (Product, error) {
	product, err := s.repository.FirstByID(ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) ListProductsByCategoryID(categoryID int) ([]Product, error) {
	productsByCategori, err := s.repository.FindByCategoryID(categoryID)
	if err != nil {
		return productsByCategori, err
	}

	return productsByCategori, nil
}
