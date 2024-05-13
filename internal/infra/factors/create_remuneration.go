package factors

import (
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/infra/external"
	controller "my-saas-app/internal/interfaces/controllers"
	"net/http"
	// "my-saas-app/internal/domain/controllers"
)

func MakeCreateRemunerationUseCase(remunerationRepository repositories.RemunerationRepository, w http.ResponseWriter, r *http.Request) {
	remunerationUseCase := usecases.NewRemunerationUseCase(remunerationRepository)
	remunerationController := controller.NewRemunerationController(remunerationUseCase)
	remunerationController.CreateRemuneration(w, r, external.GenerateJWT)
}
