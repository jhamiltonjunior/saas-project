package database

import (
	"my-saas-app/internal/domain/entities"

	"gorm.io/gorm"
)

type GormCreditCardRepositoryImpl struct {
	db *gorm.DB
}

func NewGormCreditCardRepository(db *gorm.DB) *GormCreditCardRepositoryImpl {
	return &GormCreditCardRepositoryImpl{db: db}
}

func (g *GormCreditCardRepositoryImpl) FindByID(id int32) (*entities.CreditCard, error) {
	var creditCard entities.CreditCard
	if err := g.db.First(&creditCard, id).Error; err != nil {
		return nil, err
	}
	return &creditCard, nil
}

func (g *GormCreditCardRepositoryImpl) FindByName(name string) (*entities.CreditCard, error) {
	var creditCard entities.CreditCard
	if err := g.db.Where("name = ?", name).First(&creditCard).Error; err != nil {
		return nil, err
	}
	return &creditCard, nil
}

func (g *GormCreditCardRepositoryImpl) Create(creditCard *entities.CreditCard) (int32, error) {
	creditCardDB := g.db.Create(creditCard)
	return creditCard.ID, creditCardDB.Error
}

func (g *GormCreditCardRepositoryImpl) Update(creditCard *entities.CreditCard) error {
	return g.db.Save(creditCard).Error
}

func (g *GormCreditCardRepositoryImpl) Delete(creditCard *entities.CreditCard) error {
	return g.db.Delete(&entities.CreditCard{}, creditCard.ID).Error
}

func (g *GormCreditCardRepositoryImpl) FindAll() ([]entities.CreditCard, error) {
	var creditCards []entities.CreditCard
	if err := g.db.Find(&creditCards).Error; err != nil {
		return nil, err
	}
	return creditCards, nil
}
