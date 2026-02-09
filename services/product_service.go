package services

import (
	"errors"
	"go_kasir_api/models"
	"go_kasir_api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}



func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name) // Pass name to repository
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}

// add validation
func (s *ProductService) Create(p *models.Product) error {
    if p.Name == "" {
        return errors.New("name is required")
    }
    if p.Price <= 0 {
        return errors.New("price must be greater than 0")
    }
    if p.Stock < 0 {
        return errors.New("stock cannot be negative")
    }
    return s.repo.Create(p)
}
