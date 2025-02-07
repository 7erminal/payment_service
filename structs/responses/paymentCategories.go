package responses

import "payment_service/models"

type PaymentCategoryResponseDTO struct {
	StatusCode      int
	PaymentCategory *models.Payment_categories
	StatusDesc      string
}

type PaymentCategoriesResponseDTO struct {
	StatusCode        int
	PaymentCategories *[]interface{}
	StatusDesc        string
}
