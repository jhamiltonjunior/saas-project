package controller

import (
	"encoding/json"
	"fmt"
	"my-saas-app/src/domain/entities"
	"my-saas-app/src/external/logs"
	"my-saas-app/src/usecases"
	"net/http"
	"strconv"
)

type ContextCreditCard interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
}

// type GenerateJWT func(int) (string, error)

type CreditCardController struct {
	CreditCardUseCase *usecases.CreditCardUseCase
}

func NewCreditCardController(CreditCardUseCase *usecases.CreditCardUseCase) *CreditCardController {
	return &CreditCardController{CreditCardUseCase: CreditCardUseCase}
}

func (uc *CreditCardController) GetCreditCardByID(w http.ResponseWriter, r *http.Request) {
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
			"message": "Invalid creditCard ID",
		})
		return
	}
	fmt.Println(id)

	id2 := int32(3)

	creditCard, err := uc.CreditCardUseCase.GetCreditCardByID(id2)
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "CreditCard not found",
		})
		if err != nil {
			go fileLogger.Log(err.Error())
		}
		return
	}

	json.NewEncoder(w).Encode(creditCard)
}

func (uc *CreditCardController) CreateCreditCard(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var creditCard entities.CreditCard
	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
		})
		return
	}

	creditCardId, err := uc.CreditCardUseCase.Create(&creditCard)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	creditCardCreated, err := uc.CreditCardUseCase.GetCreditCardByID(creditCardId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to get data",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"data":    creditCardCreated,
		"message": "creditCard created",
	})
}

// func (uc *CreditCardController) UpdateCreditCard(w http.ResponseWriter, r *http.Request) {
//     var creditCard entities.CreditCard
//     if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
//         return
//     }

//     if err := uc.CreditCardUseCase.UpdateCreditCard(&creditCard); err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create creditCard"})
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(creditCard)
// }

// func (uc *CreditCardController) DeleteCreditCard(w http.ResponseWriter, r *http.Request) {
//     fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
//     if err != nil {
//         panic(err)
//     }
//     defer fileLogger.Close()

//     id, err := strconv.Atoi(r.URL.Query().Get("id"))
//     if err != nil {
//         go fileLogger.Log(err.Error())
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid creditCard ID"})
//         return
//     }

//     if err := uc.CreditCardUseCase.DeleteCreditCard(id); err != nil {
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete creditCard"})
//         return
//     }

//     json.NewEncoder(w).Encode(map[string]string{"message": "CreditCard deleted successfully"})
// }
