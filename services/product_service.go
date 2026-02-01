package services

import (
	"go-crud/models"
	"go-crud/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) Update(id int, product *models.Product) error {
	product.ID = id
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
