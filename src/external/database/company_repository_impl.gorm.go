package database

import (
	"my-saas-app/src/domain/entities"

	"gorm.io/gorm"
)

type GormCompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewGormCompanyRepository(db *gorm.DB) *GormCompanyRepositoryImpl {
	return &GormCompanyRepositoryImpl{db: db}
}

func (g *GormCompanyRepositoryImpl) FindByID(id int) (*entities.Company, error) {
	var company entities.Company
	if err := g.db.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (g *GormCompanyRepositoryImpl) FindByCNPJ(cnpj string) (*entities.Company, error) {
	var company entities.Company
	if err := g.db.Where("cnpj = ?", cnpj).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (g *GormCompanyRepositoryImpl) Create(company *entities.Company) (int, error) {
	companyDB := g.db.Create(company)
	return company.ID, companyDB.Error
}

func (g *GormCompanyRepositoryImpl) Update(company *entities.Company) error {
	return g.db.Save(company).Error
}

func (g *GormCompanyRepositoryImpl) Delete(id int) error {
	return g.db.Delete(&entities.Company{}, id).Error
}
