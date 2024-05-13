package controller

import (
	"encoding/json"
	"fmt"
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/infra/logs"
	"net/http"
	"strconv"
)

type ContextRemuneration interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
}

// type GenerateJWT func(int) (string, error)

type RemunerationController struct {
	RemunerationUseCase *usecases.RemunerationUseCase
}

func NewRemunerationController(RemunerationUseCase *usecases.RemunerationUseCase) *RemunerationController {
	return &RemunerationController{RemunerationUseCase: RemunerationUseCase}
}

func (uc *RemunerationController) GetRemunerationByID(w http.ResponseWriter, r *http.Request) {
	fileLogger, err := logs.NewFileLogger("controllers_error.log")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
	}
	defer fileLogger.Close()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid remuneration ID",
		})
		return
	}

	remuneration, err := uc.RemunerationUseCase.GetRemunerationByID(id)
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Remuneration not found",
		})
		if err != nil {
			go fileLogger.Log(err.Error())
		}
		return
	}

	json.NewEncoder(w).Encode(remuneration)
}

func (uc *RemunerationController) CreateRemuneration(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var remuneration entities.Remuneration
	if err := json.NewDecoder(r.Body).Decode(&remuneration); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
		})
		return
	}

	remunerationId, err := uc.RemunerationUseCase.Create(&remuneration)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	remunerationCreated, err := uc.RemunerationUseCase.GetRemunerationByID(remunerationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to get data",
		})
		return
	}

	token, err := jwt(remunerationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"data":    remunerationCreated,
		"message": "remuneration created",
		"token":   token,
	})
}

func (uc *RemunerationController) GetAll(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	// var remuneration
	var period repositories.Period

	if err := json.NewDecoder(r.Body).Decode(&period); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
			"period":  fmt.Sprintf("%v", period),
		})
		return
	}

	remunerations, err := uc.RemunerationUseCase.GetAllRemunerationByMonths(&period)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"data":    remunerations,
		"message": "remuneration created",
		// "token":  token,
	})
}

func (uc *RemunerationController) GetAllByYear(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var year repositories.OnlyYearPeriod

	if err := json.NewDecoder(r.Body).Decode(&year); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
			"year":  fmt.Sprintf("%v", year),
		})
		return
	}

	remunerations, err := uc.RemunerationUseCase.GetAllRemunerationByYear(&year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"data":    remunerations,
		"message": "remuneration created",
		// "token":  token,
	})
}

// func (uc *RemunerationController) DeleteRemuneration(w http.ResponseWriter, r *http.Request) {
//     fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
//     if err != nil {
//         panic(err)
//     }
//     defer fileLogger.Close()

//     id, err := strconv.Atoi(r.URL.Query().Get("id"))
//     if err != nil {
//         go fileLogger.Log(err.Error())
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid remuneration ID"})
//         return
//     }

//     if err := uc.RemunerationUseCase.DeleteRemuneration(id); err != nil {
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete remuneration"})
//         return
//     }

//     json.NewEncoder(w).Encode(map[string]string{"message": "Remuneration deleted successfully"})
// }
