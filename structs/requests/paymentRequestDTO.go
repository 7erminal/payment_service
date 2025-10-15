package requests

import "payment_service/models"

type PaymentRequestDTO struct {
	InitiatedBy     int64
	Amount          float32
	Service         int64
	Sender          int64
	Reciever        int64
	PaymentMethod   int64
	TransactionId   int64
	PaymentProofUrl string
	ReferenceNumber string
}

type PaymentRequest2DTO struct {
	InitiatedBy     int64
	Amount          float32
	Service         int64
	Sender          int64
	Reciever        int64
	PaymentMethod   int64
	TransactionId   int64
	PaymentProofUrl string
	ReferenceNumber string
	CallThirdParty  bool
	Operator        string
}

type MomoPaymentRequestDTO struct {
	Payment            models.Payments
	CustomerName       string
	CustomerMsisdn     string
	CustomerEmail      string
	Channel            string
	Amount             float32
	PrimaryCallbackUrl string
	Description        string
	ClientReference    string
}

type HubtelMomoPaymentRequestDTO struct {
	CustomerName       string
	CustomerMsisdn     string
	CustomerEmail      string
	Channel            string
	Amount             float32
	PrimaryCallbackUrl string
	Description        string
	ClientReference    string
	Network            string
}

type NameInquiryRequestDTO struct {
	CustomerMsisdn string
	Channel        string
}

type HubtelNameInquiryRequestDTO struct {
	CustomerMsisdn string
	Channel        string
}
