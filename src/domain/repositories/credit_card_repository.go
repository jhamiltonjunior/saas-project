package repositories

import (
	"my-saas-app/src/domain/entities"
)

type CreditCardRepository interface {
	FindByID(id int32) (*entities.CreditCard, error)
	FindByName(name string) (*entities.CreditCard, error)
	FindAll() ([]entities.CreditCard, error)
	Create(creditCard *entities.CreditCard) (int32, error)
	Update(creditCard *entities.CreditCard) error
	Delete(creditCard *entities.CreditCard) error
}
