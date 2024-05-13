package factors

import (
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/infra/external"
	controller "my-saas-app/internal/interfaces/controllers"
	"net/http"
	// "my-saas-app/internal/domain/controllers"
)

func MakeCreateCreditCardUseCase(creditCardRepository repositories.CreditCardRepository, w http.ResponseWriter, r *http.Request) {
	creditCardUseCase := usecases.NewCreditCardUseCase(creditCardRepository)
	creditCardController := controller.NewCreditCardController(creditCardUseCase)
	creditCardController.CreateCreditCard(w, r, external.GenerateJWT)
}
