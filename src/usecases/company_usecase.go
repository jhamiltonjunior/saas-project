package usecases

import (
	"my-saas-app/src/domain/entities"
	"my-saas-app/src/domain/repositories"
)

type CompanyUseCase struct {
	companyRepository repositories.CompanyRepository
}

func NewCompanyUseCase(companyRepository repositories.CompanyRepository) *CompanyUseCase {
	return &CompanyUseCase{
		companyRepository: companyRepository,
	}
}

func (uc *CompanyUseCase) Create(input *entities.Company) (int, error) {
	company := &entities.Company{
		Name: input.Name,
		CNPJ: input.CNPJ,
	}

	companyId, err := uc.companyRepository.Create(company)

	if err != nil {
		return 0, err
	}

	return companyId, nil
}

func (uc *CompanyUseCase) GetCompanyByID(id int) (*entities.Company, error) {
	return uc.companyRepository.FindByID(id)
}

func (uc *CompanyUseCase) GetCompanyByCNPJ(cnpj string) (*entities.Company, error) {

	return uc.companyRepository.FindByCNPJ(cnpj)
}

func (uc *CompanyUseCase) Update(input *entities.Company) error {
	company := &entities.Company{
		ID:   input.ID,
		Name: input.Name,
		CNPJ: input.CNPJ,
	}

	return uc.companyRepository.Update(company)
}

func (uc *CompanyUseCase) Delete(id int) error {
	return uc.companyRepository.Delete(id)
}
