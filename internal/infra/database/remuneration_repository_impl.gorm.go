package database

import (
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormRemunerationRepositoryImpl struct {
	db *gorm.DB
}

func NewGormRemunerationRepository(db *gorm.DB) *GormRemunerationRepositoryImpl {
	return &GormRemunerationRepositoryImpl{db: db}
}

func (g *GormRemunerationRepositoryImpl) FindByID(id int) (*entities.Remuneration, error) {
	var remuneration entities.Remuneration
	if err := g.db.First(&remuneration, id).Error; err != nil {
		return nil, err
	}
	return &remuneration, nil
}

func (g *GormRemunerationRepositoryImpl) FindByName(name string) (*entities.Remuneration, error) {
	var remuneration entities.Remuneration

	name = "%" + name + "%"

	if err := g.db.Where("name Like ?", name).Limit(10).Find(&remuneration).Error; err != nil {
		return nil, err
	}
	return &remuneration, nil
}

func (g *GormRemunerationRepositoryImpl) Create(remuneration *entities.Remuneration) (int, error) {
	remunerationDB := g.db.Create(remuneration)
	return remuneration.ID, remunerationDB.Error
}

func (g *GormRemunerationRepositoryImpl) Update(remuneration *entities.Remuneration) error {
	return g.db.Save(remuneration).Error
}

func (g *GormRemunerationRepositoryImpl) Delete(remuneration *entities.Remuneration) error {
	return g.db.Delete(&entities.Remuneration{}, remuneration.ID).Error
}

func (g *GormRemunerationRepositoryImpl) FindAll(period *repositories.Period) ([]entities.Remuneration, error) {
	var remunerations []entities.Remuneration
	if err := g.db.Where("YEAR(create_at) = ? AND MONTH(create_at) = ?", period.Year, period.Month).Find(&remunerations).Error; err != nil {
		return nil, err
	}
	return remunerations, nil
}

func (g *GormRemunerationRepositoryImpl) FindAllByYear(period *repositories.OnlyYearPeriod) ([]entities.Remuneration, error) {
	var remunerations []entities.Remuneration
	if err := g.db.Where("YEAR(create_at) = ?", period.Year).Order("create_at asc").Find(&remunerations).Error; err != nil {
		return nil, err
	}
	return remunerations, nil
}

func (g *GormRemunerationRepositoryImpl) FindByType(typeRemuneration string) (*entities.Remuneration, error) {
	var remuneration entities.Remuneration
	if err := g.db.Where("type = ?", typeRemuneration).First(&remuneration).Error; err != nil {
		return nil, err
	}
	return &remuneration, nil
}
