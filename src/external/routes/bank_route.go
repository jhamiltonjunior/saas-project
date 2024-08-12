package routes

import (
	"encoding/json"
	"fmt"
	"my-saas-app/src/external/database"
	"my-saas-app/src/external/factors"
	"my-saas-app/src/external/logs"
	middlewares "my-saas-app/src/external/middleware"
	"net/http"
)

func CreateBank(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(w, r)

	switch r.Method {
	case http.MethodPost:
		middlewares.EnableCors(w, r)

		fileLogger, err := logs.NewFileLogger("general.log")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// go fileLogger.Log(fmt.Sprintf("Internal server error: %s", err.Error()))
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("Internal server error: %s", err.Error()),
			})
			return
		}
		defer fileLogger.Close()

		gorm, err := database.NewGormConnection(NewRouteVariableDB())
		if err != nil {
			go fileLogger.Log(fmt.Sprintf("Internal server error: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Internal server error",
			})
			return
		}

		factors.MakeCreateBankUseCase(database.NewGormBankRepository(gorm), w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request method",
		})
	}
}
