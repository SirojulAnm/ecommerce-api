package product

import "gorm.io/gorm"

type Repository interface {
	Save(product Product) (Product, error)
	FirstByID(ID int) (Product, error)
	FindByCategoryID(categoryID int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FirstByID(ID int) (Product, error) {
	var product Product

	err := r.db.First(&product, ID).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByCategoryID(categoryID int) ([]Product, error) {
	var product []Product

	err := r.db.Where("category_id = ?", categoryID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
