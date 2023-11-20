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

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		middlewares.EnableCors(w, r)
		fileLogger, err := logs.NewFileLogger("general.log")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Internal server error: %s", err.Error())})
			return
		}
		defer fileLogger.Close()

		gorm, err := database.NewGormConnection(NewRouteVariable())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			go fileLogger.Log(fmt.Sprintf("Internal server error: %s", err.Error()))
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}

		factors.MakeCreateCompanyUseCase(database.NewGormCompanyRepository(gorm), w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
	}
}
