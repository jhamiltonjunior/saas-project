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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	json.NewEncoder(w).Encode("Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", routes.CreateUser)
	http.ListenAndServe(":3001", nil)
}