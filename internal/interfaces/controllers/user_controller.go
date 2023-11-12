package controller

import (
	"encoding/json"
	"my-saas-app/internal/application/services"
	"my-saas-app/internal/domain/entities"
	"my-saas-app/internal/infrastructure/logging"
	"net/http"
	"strconv"
)

type Context interface {
    Param(string) string
    Bind(interface{}) error
    JSON(int, interface{})
}

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
    fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
    if err != nil {
        panic(err)
    }
    defer fileLogger.Close()

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        go fileLogger.Log(err.Error())
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
        return
    }

    user, err := uc.userService.GetUserByID(id)
    if err != nil {
        go fileLogger.Log(err.Error())
        w.WriteHeader(http.StatusNotFound)
        err := json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
        if err != nil {
            go fileLogger.Log(err.Error())
        }
        return
    }

    json.NewEncoder(w).Encode(user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user entities.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
        return
    }

    if err := uc.userService.CreateUser(&user); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create user"})
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
    var user entities.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
        return
    }

    if err := uc.userService.UpdateUser(&user); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create user"})
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
    fileLogger, err := logging.NewFileLogger("../../infrastructure/logging/logs/controllers_error.log")
    if err != nil {
        panic(err)
    }
    defer fileLogger.Close()
    
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        go fileLogger.Log(err.Error())
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
        return
    }

    if err := uc.userService.DeleteUser(id); err != nil {
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete user"})
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}