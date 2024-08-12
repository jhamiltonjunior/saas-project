package factors

import (
	controller "my-saas-app/src/adapters/controllers"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/external/external"
	"my-saas-app/src/usecases"
	"net/http"
	// "my-saas-app/src/domain/controllers"
)

func MakeCreateUserUseCase(userRepository repositories.UserRepository, w http.ResponseWriter, r *http.Request) {
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	userController.CreateUser(w, r, external.GenerateJWT)
}
