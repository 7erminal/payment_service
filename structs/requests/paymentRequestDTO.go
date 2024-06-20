package requests

type PaymentRequestDTO struct {
	InitiatedBy   int64
	Amount        float32
	Service       int64
	Sender        int64
	Reciever      int64
	PaymentMethod int64
	TransactionId int64
}
