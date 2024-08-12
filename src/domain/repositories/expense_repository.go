package repositories

import (
	"my-saas-app/src/domain/entities"
)

// type Period struct {
// 	Month int `json:"month"`
// 	Year  int `json:"year"`
// }

// type OnlyYearPeriod struct {
// 	Year  int `json:"year"`
// }

type ExpenseRepository interface {
	FindByID(id *entities.ExpenseID) (*entities.Expense, error)
	FindAll(*Period) ([]entities.Expense, error)
	FindAllByYear(*OnlyYearPeriod) ([]entities.Expense, error)
	FindByName(cnpj string) (*entities.Expense, error)
	FindByType(cnpj string) (*entities.Expense, error)
	Create(expense *entities.Expense) (entities.ExpenseID, error)
	Update(expense *entities.Expense) error
	Delete(expense *entities.Expense) error
}
