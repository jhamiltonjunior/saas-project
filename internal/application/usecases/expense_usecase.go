package usecases

import (
	"errors"
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"
)

type ExpenseUseCase struct {
	expenseRepository repositories.ExpenseRepository
}

func NewExpenseUseCase(expenseRepository repositories.ExpenseRepository) *ExpenseUseCase {
	return &ExpenseUseCase{
		expenseRepository: expenseRepository,
	}
}

func (bk *ExpenseUseCase) Create(input *entities.Expense) (entities.ExpenseID, error) {
	expense := &entities.Expense{
		Name:         input.Name,
		Value:        input.Value,
		UserID:       input.UserID,
		RecurrenceID: input.RecurrenceID,
	}

	expenseId, err := bk.expenseRepository.Create(expense)

	if err != nil {
		return 0, err
	}

	return expenseId, nil
}

func (bk *ExpenseUseCase) GetExpenseByID(id *entities.ExpenseID) (*entities.Expense, error) {
	return bk.expenseRepository.FindByID(id)
}

func (bk *ExpenseUseCase) GetExpenseByName(name string) (*entities.Expense, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return bk.expenseRepository.FindByName(name)
}

func (bk *ExpenseUseCase) GetAllExpenseByMonths(period *repositories.Period) ([]entities.Expense, error) {
	return bk.expenseRepository.FindAll(period)
}

func (bk *ExpenseUseCase) GetAllExpenseByYear(year *repositories.OnlyYearPeriod) ([]entities.Expense, error) {
	return bk.expenseRepository.FindAllByYear(year)
}

func (bk *ExpenseUseCase) Update(input *entities.Expense) error {
	expense := &entities.Expense{
		ID:     input.ID,
		Name:   input.Name,
		Value:  input.Value,
		UserID: input.UserID,
	}

	return bk.expenseRepository.Update(expense)
}

func (bk *ExpenseUseCase) Delete(expense *entities.Expense) error {
	return bk.expenseRepository.Delete(expense)
}
