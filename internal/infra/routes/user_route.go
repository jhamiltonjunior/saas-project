package routes

import (
	"encoding/json"
	"fmt"
	"my-saas-app/internal/infra/database"
	"my-saas-app/internal/infra/factors"
	"my-saas-app/internal/infra/logs"
	middlewares "my-saas-app/internal/infra/middleware"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		middlewares.EnableCors(w, r)
		fileLogger, err := logs.NewFileLogger("general.log")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// go fileLogger.Log(fmt.Sprintf("Internal server error: %s", err.Error()))
			json.NewEncoder(w).Encode(map[string]string{
				"status": "error",
				"message": fmt.Sprintf("Internal server error: %s", err.Error()),
			})
			return
		}
		defer fileLogger.Close()

		gorm, err := database.NewGormConnection("root:0000@tcp(localhost:3306)/my_saas_app?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			go fileLogger.Log(fmt.Sprintf("Internal server error: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "error",
				"message": "Internal server error",
			})
			return
		}

		factors.MakeCreateUserUseCase(database.NewGormUserRepository(gorm), w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Invalid request method",
		})
	}
}
