package controller

import (
	"encoding/json"
	"my-saas-app/src/domain/entities"
	"my-saas-app/src/external/logs"
	"my-saas-app/src/usecases"
	"net/http"
	"strconv"
)

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
}

type GenerateJWT func(int) (string, error)

type UserController struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserController(UserUseCase *usecases.UserUseCase) *UserController {
	return &UserController{UserUseCase: UserUseCase}
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
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
			"message": "Invalid user ID",
		})
		return
	}

	user, err := uc.UserUseCase.GetUserByID(id)
	if err != nil {
		go fileLogger.Log(err.Error())
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "User not found",
		})
		if err != nil {
			go fileLogger.Log(err.Error())
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, jwt GenerateJWT) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request payload",
		})
		return
	}

	userId, err := uc.UserUseCase.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	userCreated, err := uc.UserUseCase.GetUserByID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to get data",
		})
		return
	}

	token, err := jwt(userId)
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
		"data":    userCreated,
		"message": "user created",
		"token":   token,
	})
}

// func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
//     var user entities.User
//     if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
//         return
//     }

//     if err := uc.UserUseCase.UpdateUser(&user); err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create user"})
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(user)
// }

// func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
//     fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
//     if err != nil {
//         panic(err)
//     }
//     defer fileLogger.Close()

//     id, err := strconv.Atoi(r.URL.Query().Get("id"))
//     if err != nil {
//         go fileLogger.Log(err.Error())
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
//         return
//     }

//     if err := uc.UserUseCase.DeleteUser(id); err != nil {
//         json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete user"})
//         return
//     }

//     json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
// }
