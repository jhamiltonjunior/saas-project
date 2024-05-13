package repositories

import (
	"my-saas-app/internal/domain/entities"
)

type BankRepository interface {
	FindByID(id int) (*entities.Bank, error)
	FindAll() ([]entities.Bank, error)
	FindByName(name string) (*entities.Bank, error)
	Create(bank *entities.Bank) (int, error)
	Update(bank *entities.Bank) error
	Delete(id int) error
}