package factors

import (
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/interfaces/controllers"
	"net/http"
)

func MakeCreateCompanyUseCase(companyRepository repositories.CompanyRepository, w http.ResponseWriter, r *http.Request) {
	companyUseCase := usecases.NewCompanyUseCase(companyRepository)
	company := controller.NewCompanyController(companyUseCase)
	company.CreateCompany(w, r)
}