package responses

import (
	"payment_service/models"
	"time"
)

type PaymentResponse struct {
	Sender          string
	Reciever        string
	Amount          float64
	Commission      float64
	Charge          float64
	OtherCharge     float64
	PaymentAmount   float64
	PaymentMethod   *models.Payment_methods
	PaymentProof    string
	Status          *models.Status
	Service         string
	SenderAccount   string
	ReceiverAccount string
	ReferenceNumber string
	DateCreated     time.Time `orm:"type(datetime)"`
	DateModified    time.Time `orm:"type(datetime)"`
	ProcessedDate   time.Time `orm:"type(datetime);null"`
	Active          int
	CallbackUrl     string
}

type PaymentResponseDTO struct {
	StatusCode int
	Payment    *PaymentResponse
	StatusDesc string
}

type PaymentsResponseDTO struct {
	StatusCode int
	Payments   *[]interface{}
	StatusDesc string
}

type RequestMoneyDataResponse struct {
	PaymentId          string
	Description        string
	Amount             float64
	Charges            float64
	SenderAccount      string
	ReceiverAccount    string
	AmountAfterCharges float64
	AmountCharged      float64
	ReferenceNumber    string
	PaymentDate        string
}

type RequestMoneyResponseDTO struct {
	StatusCode int
	Result     *RequestMoneyDataResponse
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
	Amount                float64
	Charges               float64
	AmountAfterCharges    float64
	AmountCharged         float64
	ClientReference       string
	DeliveryFee           float64
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
