package controller

import (
	"encoding/json"
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/infra/logs"
	"net/http"
	"strconv"
)

type ContextBank interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
}

// type GenerateJWT func(int) (string, error)

type BankController struct {
	BankUseCase *usecases.BankUseCase
}

func NewBankController(BankUseCase *usecases.BankUseCase) *BankController {
	return &BankController{BankUseCase: BankUseCase}
}

func (uc *BankController) GetBankByID(w http.ResponseWriter, r *http.Request) {
	fileLogger, err := logs.NewFileLogger("controllers_error.log")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Internal Server Error",
		})
	}
	defer fileLogger.Close()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Invalid bank ID",
		})
		return
	}

	bank, err := uc.BankUseCase.GetBankByID(id)
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Bank not found",
		})
		if err != nil {
			go fileLogger.Log(err.Error())
		}
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func (uc *BankController) CreateBank(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var bank entities.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Invalid request payload",
		})
		return
	}

	bankId, err := uc.BankUseCase.Create(&bank)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	bankCreated, err := uc.BankUseCase.GetBankByID(bankId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Failed to get data",
		})
		return
	}

	token, err := jwt(bankId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"message": "Failed to generate token",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   bankCreated,
		"message":    "bank created",
		"token":  token,
	})
}

// func (uc *BankController) UpdateBank(w http.ResponseWriter, r *http.Request) {
//     var bank entities.Bank
//     if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
//         return
//     }

//     if err := uc.BankUseCase.UpdateBank(&bank); err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create bank"})
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(bank)
// }

// func (uc *BankController) DeleteBank(w http.ResponseWriter, r *http.Request) {
//     fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
//     if err != nil {
//         panic(err)
//     }
//     defer fileLogger.Close()

//     id, err := strconv.Atoi(r.URL.Query().Get("id"))
//     if err != nil {
//         go fileLogger.Log(err.Error())
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid bank ID"})
//         return
//     }

//     if err := uc.BankUseCase.DeleteBank(id); err != nil {
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete bank"})
//         return
//     }

//     json.NewEncoder(w).Encode(map[string]string{"message": "Bank deleted successfully"})
// }
