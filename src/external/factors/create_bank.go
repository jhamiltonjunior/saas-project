package factors

import (
	controller "my-saas-app/src/adapters/controllers"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/external/external"
	"my-saas-app/src/usecases"
	"net/http"
	// "my-saas-app/src/domain/controllers"
)

func MakeCreateBankUseCase(bankRepository repositories.BankRepository, w http.ResponseWriter, r *http.Request) {
	bankUseCase := usecases.NewBankUseCase(bankRepository)
	bankController := controller.NewBankController(bankUseCase)
	bankController.CreateBank(w, r, external.GenerateJWT)
}
