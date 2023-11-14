package routes

import (
	"encoding/json"
	"fmt"
	"my-saas-app/internal/infra/database"
	"my-saas-app/internal/infra/factors"
	"my-saas-app/internal/infra/logs"
	"my-saas-app/internal/infra/middleware"
	"net/http"
)

// tente testar o erro que da quando instancia o gorm
// pois se eu n dar return ele vai continuar executando o codigo

func CreateUser(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(w)
	
	switch r.Method {
	case http.MethodPost:
		fileLogger, err := logs.NewFileLogger("general.log")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Internal server error: %s", err.Error())})
			return
		}
		defer fileLogger.Close()

		gorm, err := database.NewGormConnection("root:0000@tcp(localhost:3306)/my_sas_app?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			fileLogger.Log("Internal server error")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}

		factors.MakeCreateUserUseCase(database.NewGormRepository(gorm), w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
	}
}
