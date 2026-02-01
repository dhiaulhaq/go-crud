package repositories

import (
	"database/sql"
	"go-crud/models"
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

func (r *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
}

func (r *CategoryRepository) Update(category *models.Category) error {
	query := "UPDATE categories SET name=$1, description=$2 WHERE id=$3"
	_, err := r.db.Exec(query, category.Name, category.Description, category.ID)
	return err
}

func (r *CategoryRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id=$1", id)
	return err
}
