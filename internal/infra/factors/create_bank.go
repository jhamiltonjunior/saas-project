package factors

import (
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/infra/external"
	controller "my-saas-app/internal/interfaces/controllers"
	"net/http"
	// "my-saas-app/internal/domain/controllers"
)

func MakeCreateBankUseCase(bankRepository repositories.BankRepository, w http.ResponseWriter, r *http.Request) {
	bankUseCase := usecases.NewBankUseCase(bankRepository)
	bankController := controller.NewBankController(bankUseCase)
	bankController.CreateBank(w, r, external.GenerateJWT)
}
