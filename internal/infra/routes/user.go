package routes

import (
	"encoding/json"
	"my-saas-app/internal/infra/database"
	"my-saas-app/internal/infra/factors"
	middlewares "my-saas-app/internal/infra/middleware"
	"my-saas-app/internal/infra/logging"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		middlewares.EnableCors(w)
		gorm, err := database.NewGormConnection("root:0000@tcp(localhost:3306)/my_saas_app?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
			go fileLogger.Log(err.Error())
		}
		factors.MakeCreateUserUseCase(database.NewGormRepository(gorm), w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
	}
}
