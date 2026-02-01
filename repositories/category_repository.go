package repositories

import (
	"go_kasir_api/models"
)

var categories = []models.Category{
	{ID: 1, Name: "Makanan", Description: "Produk makanan"},
	{ID: 2, Name: "Minuman", Description: "Produk minuman"},
}

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) GetAll() []models.Category {
	return categories
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	category.ID = len(categories) + 1
	categories = append(categories, *category)
	return nil
}

func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	for _, c := range categories {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, nil
}

func (repo *CategoryRepository) Update(category *models.Category) error {
	for i := range categories {
		if categories[i].ID == category.ID {
			categories[i] = *category
			return nil
		}
	}
	return nil
}

func (repo *CategoryRepository) Delete(id int) error {
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}
	return nil
}
