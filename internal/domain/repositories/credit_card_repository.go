package repositories

import (
	"my-saas-app/internal/domain/entities"
)

type CreditCardRepository interface {
	FindByID(id int) (*entities.CreditCard, error)
	FindByName(name string) (*entities.CreditCard, error)
	FindAll() ([]entities.CreditCard, error)
	Create(creditCard *entities.CreditCard) (int, error)
	Update(creditCard *entities.CreditCard) error
	Delete(creditCard *entities.CreditCard) error
}