package repositories

import (
	"my-saas-app/src/domain/entities"
)

type CompanyRepository interface {
	FindByID(id int) (*entities.Company, error)
	FindByCNPJ(cnpj string) (*entities.Company, error)
	Create(company *entities.Company) (int, error)
	Update(company *entities.Company) error
	Delete(id int) error
}
