package requests

type ExpenseRequest struct {
	Currency      int64
	Category      int64
	Description   string
	Amount        float64
	Date          string
	PaymentMethod int64
}
