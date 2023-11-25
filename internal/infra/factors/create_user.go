package factors

import (
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/infra/external"
	controller "my-saas-app/internal/interfaces/controllers"
	"net/http"
	// "my-saas-app/internal/domain/controllers"
)

func MakeCreateUserUseCase(userRepository repositories.UserRepository, w http.ResponseWriter, r *http.Request) {
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	userController.CreateUser(w, r, external.GenerateJWT)
}
