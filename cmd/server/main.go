package main

import (
	"encoding/json"
	"my-saas-app/internal/infra/routes"
	"net/http"
	// "github.com/my-saas-app/internal/interfaces/api"
	// "github.com/my-saas-app/internal/interfaces/config"
	// "github.com/my-saas-app/internal/interfaces/routes"
	// "github.com/my-saas-app/pkg/server"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	json.NewEncoder(w).Encode("Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/user", routes.CreateUser)
	http.HandleFunc("/api/company", routes.CreateCompany)
	http.ListenAndServe(":3000", nil)
}