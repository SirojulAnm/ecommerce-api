package category

type Service interface {
	AddCategory(input CategoryInput) (Category, error)
	GetAll() ([]Category, error)
	GetCategoryByID(ID int) (Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddCategory(input CategoryInput) (Category, error) {
	category := Category{}
	category.Name = input.Name
	category.Description = input.Description

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) GetAll() ([]Category, error) {
	categorys, err := s.repository.GetAll()
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}

func (s *service) GetCategoryByID(ID int) (Category, error) {
	categorys, err := s.repository.FirstByID(ID)
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}
