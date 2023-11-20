package controller

import (
	"encoding/json"
	"net/http"

	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/entities"
)

type CompanyController struct {
	companyUseCase *usecases.CompanyUseCase
}

func NewCompanyController(companyUseCase *usecases.CompanyUseCase) *CompanyController {
	return &CompanyController{companyUseCase: companyUseCase}
}

func (cc *CompanyController) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company entities.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	companyId, err := cc.companyUseCase.Create(&company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error creating company"})
		return
	}

	companyCreated, err := cc.companyUseCase.GetCompanyByID(companyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error to get your data"})
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(companyCreated)
}