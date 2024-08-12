package factors

import (
	controller "my-saas-app/src/adapters/controllers"
	"my-saas-app/src/domain/repositories"
	"my-saas-app/src/external/external"
	"my-saas-app/src/usecases"
	"net/http"
	// "my-saas-app/src/domain/controllers"
)

func MakeCreateExpenseUseCase(expenseRepository repositories.ExpenseRepository, w http.ResponseWriter, r *http.Request) {
	expenseUseCase := usecases.NewExpenseUseCase(expenseRepository)
	expenseController := controller.NewExpenseController(expenseUseCase)
	expenseController.CreateExpense(w, r, external.GenerateJWT)
}
