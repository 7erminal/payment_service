package responses

import "payment_service/models"

type ExpenseResponse struct {
	StatusCode int
	Expense    *models.Expense_records
	StatusDesc string
}

type ExpensesResponseDTO struct {
	StatusCode int
	Expenses   *[]interface{}
	StatusDesc string
}
