package repositories

import (
	"database/sql"
	"errors"
	"go-crud/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name 
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`

	var p models.Product

	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName)

	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) GetAll(nameFilter string) ([]models.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products"
	args := []interface{}{}

	if nameFilter != "" {
		query += " WHERE name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name=$1, price=$2, stock=$3, category_id=$4 WHERE id=$5"
	_, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	return err
}

func (repo *ProductRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
