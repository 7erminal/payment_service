package responses

import (
	"payment_service/models"
	"time"
)

type PaymentResponse struct {
	PaymentId      int64 `orm:"auto"`
	InitiatedBy    int64
	Sender         *models.Users
	Reciever       *models.Users
	Amount         float64
	PaymentMethod  *models.Payment_methods
	Status         *models.Status
	PaymentAccount int
	DateCreated    time.Time `orm:"type(datetime)"`
	DateModified   time.Time `orm:"type(datetime)"`
	CreatedBy      int
	ModifiedBy     int
	Active         int
}

type PaymentResponseDTO struct {
	StatusCode int
	Payment    *models.Payments
	StatusDesc string
}
