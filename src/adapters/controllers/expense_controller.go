package controller

import (
	"encoding/json"
	"fmt"
	"my-saas-app/src/domain/entities"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/external/logs"
	"my-saas-app/src/usecases"
	"net/http"
)

type ContextExpense interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
}

// type GenerateJWT func(int) (string, error)

type ExpenseController struct {
	ExpenseUseCase *usecases.ExpenseUseCase
}

func NewExpenseController(ExpenseUseCase *usecases.ExpenseUseCase) *ExpenseController {
	return &ExpenseController{ExpenseUseCase: ExpenseUseCase}
}

func (uc *ExpenseController) GetExpenseByID(w http.ResponseWriter, r *http.Request) {
	fileLogger, err := logs.NewFileLogger("controllers_error.log")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Internal Server Error",
		})
	}
	defer fileLogger.Close()

	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil {
	// 	go fileLogger.Log(err.Error())
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]string{
	// 		"status":  "error",
	// 		"message": "Invalid expense ID",
	// 	})
	// 	return
	// }

	id := entities.ExpenseID(32)
	expense, err := uc.ExpenseUseCase.GetExpenseByID(&id)
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Expense not found",
		})
		if err != nil {
			go fileLogger.Log(err.Error())
		}
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func (uc *ExpenseController) CreateExpense(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var expense entities.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
		})
		return
	}

	expenseId, err := uc.ExpenseUseCase.Create(&expense)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	expenseCreated, err := uc.ExpenseUseCase.GetExpenseByID(&expenseId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to get data",
		})
		return
	}

	token, err := jwt(int(expenseId))
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
		"data":    expenseCreated,
		"message": "expense created",
		"token":   token,
	})
}

func (uc *ExpenseController) GetAll(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	// var expense
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

	expenses, err := uc.ExpenseUseCase.GetAllExpenseByMonths(&period)
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
		"data":    expenses,
		"message": "expense created",
		// "token":  token,
	})
}

func (uc *ExpenseController) GetAllByYear(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var year repositories.OnlyYearPeriod

	if err := json.NewDecoder(r.Body).Decode(&year); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
			"year":    fmt.Sprintf("%v", year),
		})
		return
	}

	expenses, err := uc.ExpenseUseCase.GetAllExpenseByYear(&year)
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
		"data":    expenses,
		"message": "expense created",
		// "token":  token,
	})
}

// func (uc *ExpenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
//     fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
//     if err != nil {
//         panic(err)
//     }
//     defer fileLogger.Close()

//     id, err := strconv.Atoi(r.URL.Query().Get("id"))
//     if err != nil {
//         go fileLogger.Log(err.Error())
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid expense ID"})
//         return
//     }

//     if err := uc.ExpenseUseCase.DeleteExpense(id); err != nil {
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete expense"})
//         return
//     }

//     json.NewEncoder(w).Encode(map[string]string{"message": "Expense deleted successfully"})
// }
