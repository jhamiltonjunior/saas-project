package database

import (
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormExpenseRepositoryImpl struct {
	db *gorm.DB
}

func NewGormExpenseRepository(db *gorm.DB) *GormExpenseRepositoryImpl {
	return &GormExpenseRepositoryImpl{db: db}
}

func (g *GormExpenseRepositoryImpl) FindByID(id *entities.ExpenseID) (*entities.Expense, error) {
	var expense entities.Expense
	if err := g.db.First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (g *GormExpenseRepositoryImpl) FindByName(name string) (*entities.Expense, error) {
	var expense entities.Expense

	name = "%" + name + "%"

	if err := g.db.Where("name Like ?", name).Limit(10).Find(&expense).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (g *GormExpenseRepositoryImpl) Create(expense *entities.Expense) (entities.ExpenseID, error) {
	expenseDB := g.db.Create(expense)
	return expense.ID, expenseDB.Error
}

func (g *GormExpenseRepositoryImpl) Update(expense *entities.Expense) error {
	return g.db.Save(expense).Error
}

func (g *GormExpenseRepositoryImpl) Delete(expense *entities.Expense) error {
	return g.db.Delete(&entities.Expense{}, expense.ID).Error
}

func (g *GormExpenseRepositoryImpl) FindAll(period *repositories.Period) ([]entities.Expense, error) {
	var expenses []entities.Expense
	if err := g.db.Where("YEAR(create_at) = ? AND MONTH(create_at) = ?", period.Year, period.Month).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (g *GormExpenseRepositoryImpl) FindAllByYear(period *repositories.OnlyYearPeriod) ([]entities.Expense, error) {
	var expenses []entities.Expense
	if err := g.db.Where("YEAR(create_at) = ?", period.Year).Order("create_at asc").Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (g *GormExpenseRepositoryImpl) FindByType(typeExpense string) (*entities.Expense, error) {
	var expense entities.Expense
	if err := g.db.Where("type = ?", typeExpense).First(&expense).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}
