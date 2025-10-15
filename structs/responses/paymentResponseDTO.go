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

type PaymentsResponseDTO struct {
	StatusCode int
	Payments   *[]interface{}
	StatusDesc string
}

type PaymentMethodsResponseDTO struct {
	StatusCode     int
	PaymentMethods *[]models.Payment_methods
	StatusDesc     string
}

type HubtelPaymentRequestApiResponseData struct {
	TransactionId         string
	Description           string
	Amount                float32
	Charges               float32
	AmountAfterCharges    float32
	AmountCharged         float32
	ClientReference       string
	DeliveryFee           float32
	ExternalTransactionId string
	OrderId               string
	PaymentDate           string
}

type HubtelPaymentRequestApiResponseDTO struct {
	ResponseCode string
	Data         HubtelPaymentRequestApiResponseData
	Message      string
}

type HubtelPaymentRequestResponseDTO struct {
	Success    bool
	Result     *HubtelPaymentRequestApiResponseData
	StatusDesc string
}

type HubtelNameInquiryApiResponseData struct {
	IsRegistered string `json:"isRegistered"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Profile      string `json:"profile"`
}

type HubtelNameInquiryResponseData struct {
	IsRegistered string
	Name         string
	Status       string
	Profile      string
}

type HubtelNameInquiryApiResponseDTO struct {
	ResponseCode string
	Message      string
	Label        string
	Data         HubtelNameInquiryApiResponseData `json:"data"`
}

type HubtelNameInquiryResponseDTO struct {
	Success    bool
	Result     *HubtelNameInquiryResponseData
	StatusDesc string
}
