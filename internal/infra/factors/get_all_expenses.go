package factors

import (
	"my-saas-app/internal/application/usecases"
	"my-saas-app/internal/domain/repositories"
	"my-saas-app/internal/infra/external"
	controller "my-saas-app/internal/interfaces/controllers"
	"net/http"
	// "my-saas-app/internal/domain/controllers"
)

func MakeGetAllExpenseByMonthUseCase(expenseRepository repositories.ExpenseRepository, w http.ResponseWriter, r *http.Request) {
	expenseUseCase := usecases.NewExpenseUseCase(expenseRepository)
	expenseController := controller.NewExpenseController(expenseUseCase)
	expenseController.GetAll(w, r, external.GenerateJWT)
}

func MakeGetAllExpenseByYearUseCase(expenseRepository repositories.ExpenseRepository, w http.ResponseWriter, r *http.Request) {
	expenseUseCase := usecases.NewExpenseUseCase(expenseRepository)
	expenseController := controller.NewExpenseController(expenseUseCase)
	expenseController.GetAllByYear(w, r, external.GenerateJWT)
}
