package factors

import (
	"my-saas-app/src/adapters/controllers"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/usecases"
	"net/http"
)

func MakeCreateCompanyUseCase(companyRepository repositories.CompanyRepository, w http.ResponseWriter, r *http.Request) {
	companyUseCase := usecases.NewCompanyUseCase(companyRepository)
	company := controller.NewCompanyController(companyUseCase)
	company.CreateCompany(w, r)
}
