package services

import (
	"errors" // ✅ ADD THIS IMPORT
	"go_kasir_api/models"
	"go_kasir_api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()  
}

func (s *CategoryService) Create(c *models.Category) error {
	// Validation
	if c.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(c)
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}


func (s *CategoryService) Update(c *models.Category) error {
	// Validation
	if c.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Update(c)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}