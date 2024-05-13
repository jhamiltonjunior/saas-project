package usecases

import (
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"
)

type CreditCardUseCase struct {
	creditCardRepository repositories.CreditCardRepository
}

func NewCreditCardUseCase(creditCardRepository repositories.CreditCardRepository) *CreditCardUseCase {
	return &CreditCardUseCase{
		creditCardRepository: creditCardRepository,
	}
}

func (bk *CreditCardUseCase) Create(input *entities.CreditCard) (int32, error) {
	creditCard := &entities.CreditCard{
		Name:   input.Name,
		Value:  input.Value,
		DueDate:  input.DueDate,
		UserID: input.UserID,
		FlagID: input.FlagID,
		BankID: input.BankID,
	}

	creditCardId, err := bk.creditCardRepository.Create(creditCard)

	if err != nil {
		return 0, err
	}

	return creditCardId, nil
}

func (bk *CreditCardUseCase) GetCreditCardByID(id int32) (*entities.CreditCard, error) {
	return bk.creditCardRepository.FindByID(id)
}

func (bk *CreditCardUseCase) GetCreditCardByName(name string) (*entities.CreditCard, error) {

	return bk.creditCardRepository.FindByName(name)
}

func (bk *CreditCardUseCase) GetAllCreditCards() ([]entities.CreditCard, error) {
	return bk.creditCardRepository.FindAll()
}

func (bk *CreditCardUseCase) Update(input *entities.CreditCard) error {
	creditCard := &entities.CreditCard{
		ID:     input.ID,
		Name:   input.Name,
		Value:  input.Value,
		UserID: input.UserID,
		FlagID: input.FlagID,
	}

	return bk.creditCardRepository.Update(creditCard)
}

func (bk *CreditCardUseCase) Delete(creditCard *entities.CreditCard) error {
	return bk.creditCardRepository.Delete(creditCard)
}
