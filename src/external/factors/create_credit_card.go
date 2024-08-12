package factors

import (
	controller "my-saas-app/src/adapters/controllers"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/external/external"
	"my-saas-app/src/usecases"
	"net/http"
	// "my-saas-app/src/domain/controllers"
)

func MakeCreateCreditCardUseCase(creditCardRepository repositories.CreditCardRepository, w http.ResponseWriter, r *http.Request) {
	creditCardUseCase := usecases.NewCreditCardUseCase(creditCardRepository)
	creditCardController := controller.NewCreditCardController(creditCardUseCase)
	creditCardController.CreateCreditCard(w, r, external.GenerateJWT)
}
