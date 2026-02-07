package services

import (
	"go-crud/models"
	"go-crud/repositories"
	"time"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetReport(startDate, endDate string) (*models.SalesReport, error) {
	if startDate == "" || endDate == "" {
		today := time.Now().Format("2006-01-02")
		startDate = today
		endDate = today
	}
	return s.repo.GetReport(startDate, endDate)
}
