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
	if r.URL.Path == "/" {
		err := json.NewEncoder(w).Encode("Hello World")
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("404 not found :/")
		if err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/user/register", routes.CreateUser)
	http.HandleFunc("/api/company", routes.CreateCompany)

	http.HandleFunc("/api/bank/create", routes.CreateBank)

	http.HandleFunc("/api/credit-card/create", routes.CreateCreditCard)

	http.HandleFunc("/api/remuneration/create", routes.CreateRemuneration)
	http.HandleFunc("/api/remuneration/get-all", routes.GetAllRemunerationByMonth)
	http.HandleFunc("/api/remuneration/get-all-by-year", routes.GetAllRemunerationByYear)

	http.HandleFunc("/api/expense/create", routes.CreateExpense)
	http.HandleFunc("/api/expense/get-all", routes.GetAllExpenseByMonth)
	http.HandleFunc("/api/expense/get-all-by-year", routes.GetAllExpenseByYear)

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		panic(err)
	}
}
