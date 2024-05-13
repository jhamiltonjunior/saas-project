package usecases

import (
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"
)

type BankUseCase struct {
	bankRepository repositories.BankRepository
}

func NewBankUseCase(bankRepository repositories.BankRepository) *BankUseCase {
	return &BankUseCase{
		bankRepository: bankRepository,
	}
}

func (bk *BankUseCase) Create(input *entities.Bank) (int, error) {
	bank := &entities.Bank{
		Name: input.Name,
		Image: input.Image,
	}

	bankId, err := bk.bankRepository.Create(bank)

	if err != nil {
		return 0, err
	}

	return bankId, nil
}

func (bk *BankUseCase) GetBankByID(id int) (*entities.Bank, error) {
	return bk.bankRepository.FindByID(id)
}

func (bk *BankUseCase) GetBankByName(name string) (*entities.Bank, error) {
	
	return bk.bankRepository.FindByName(name)
}

func (bk *BankUseCase) GetAllBanks() ([]entities.Bank, error) {
	return bk.bankRepository.FindAll()
}

func (bk *BankUseCase) Update(input *entities.Bank) error {
	bank := &entities.Bank{
		ID:   input.ID,
		Name: input.Name,
		Image: input.Image,
	}

	return bk.bankRepository.Update(bank)
}

func (bk *BankUseCase) Delete(id int) error {
	return bk.bankRepository.Delete(id)
}
