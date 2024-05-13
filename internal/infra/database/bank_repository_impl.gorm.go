package database

import (
	"my-saas-app/internal/domain/entities"

	"gorm.io/gorm"
)

type GormBankRepositoryImpl struct {
	db *gorm.DB
}

func NewGormBankRepository(db *gorm.DB) *GormBankRepositoryImpl {
	return &GormBankRepositoryImpl{db: db}
}

func (g *GormBankRepositoryImpl) FindByID(id int) (*entities.Bank, error) {
	var bank entities.Bank
	if err := g.db.First(&bank, id).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func (g *GormBankRepositoryImpl) FindByName(name string) (*entities.Bank, error) {
	var bank entities.Bank
	if err := g.db.Where("name = ?", name).First(&bank).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func (g *GormBankRepositoryImpl) Create(bank *entities.Bank) (int, error) {
	bankDB := g.db.Create(bank)
	return bank.ID, bankDB.Error
}

func (g *GormBankRepositoryImpl) Update(bank *entities.Bank) error {
	return g.db.Save(bank).Error
}

func (g *GormBankRepositoryImpl) Delete(id int) error {
	return g.db.Delete(&entities.Bank{}, id).Error
}

func (g *GormBankRepositoryImpl) FindAll() ([]entities.Bank, error) {
	var banks []entities.Bank
	if err := g.db.Find(&banks).Error; err != nil {
		return nil, err
	}
	return banks, nil
}
