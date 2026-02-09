package repositories

import (
	"database/sql"
	"errors"
	"go_kasir_api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) Create(c *models.Category) error {
	return r.db.QueryRow(
		"INSERT INTO categories(name, description) VALUES($1, $2) RETURNING id",
		c.Name, c.Description,
	).Scan(&c.ID)
}

func (r *CategoryRepository) GetByID(id int) (*models.Category, error) {
	c := &models.Category{}
	err := r.db.QueryRow(
		"SELECT id, name, description FROM categories WHERE id=$1",
		id,
	).Scan(&c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}
	return c, err
}

func (r *CategoryRepository) Update(c *models.Category) error {
	result, err := r.db.Exec(
		"UPDATE categories SET name=$1, description=$2 WHERE id=$3",
		c.Name, c.Description, c.ID,
	)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("category not found")
	}
	return nil
}