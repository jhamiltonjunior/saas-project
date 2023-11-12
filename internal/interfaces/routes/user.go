package routes 

import (
	"net/http"
	"encoding/json"
	"my-saas-app/internal/interfaces/middleware"
	"my-saas-app/internal/interfaces/controllers"
	"my-saas-app/internal/domain/entities"
)

func createUser (user *entities.User) {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        userController.CreateUser(w, r)
    default:
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
})
}
