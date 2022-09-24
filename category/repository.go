package category

import "gorm.io/gorm"

type Repository interface {
	Save(category Category) (Category, error)
	GetAll() ([]Category, error)
	FirstByID(ID int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) GetAll() ([]Category, error) {
	var category []Category
	err := r.db.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FirstByID(ID int) (Category, error) {
	var category Category
	err := r.db.First(&category, ID).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
