package factors

import (
	controller "my-saas-app/src/adapters/controllers"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/external/external"
	"my-saas-app/src/usecases"
	"net/http"
	// "my-saas-app/src/domain/controllers"
)

func MakeGetAllRemunerationByMonthUseCase(remunerationRepository repositories.RemunerationRepository, w http.ResponseWriter, r *http.Request) {
	remunerationUseCase := usecases.NewRemunerationUseCase(remunerationRepository)
	remunerationController := controller.NewRemunerationController(remunerationUseCase)
	remunerationController.GetAll(w, r, external.GenerateJWT)
}

func MakeGetAllRemunerationByYearUseCase(remunerationRepository repositories.RemunerationRepository, w http.ResponseWriter, r *http.Request) {
	remunerationUseCase := usecases.NewRemunerationUseCase(remunerationRepository)
	remunerationController := controller.NewRemunerationController(remunerationUseCase)
	remunerationController.GetAllByYear(w, r, external.GenerateJWT)
}
