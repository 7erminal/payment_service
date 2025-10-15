package functions

import (
	"payment_service/models"
	"payment_service/structs/requests"
	"payment_service/structs/responses"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

func PaymentRequestViaMobileMoney(c *beego.Controller, req requests.MomoPaymentRequestDTO) (responses.HubtelPaymentRequestResponseDTO, error) {
	//do something

	momoRequest := requests.HubtelMomoPaymentRequestDTO{
		CustomerName:       req.CustomerName,
		CustomerMsisdn:     req.CustomerMsisdn,
		CustomerEmail:      req.CustomerEmail,
		Channel:            req.Channel,
		Amount:             float32(req.Amount),
		PrimaryCallbackUrl: req.PrimaryCallbackUrl,
		Description:        "Payment for " + req.CustomerName,
		ClientReference:    req.ClientReference,
		Network:            req.Channel,
	}

	responseCode := false
	responseMessage := "Error processing request"
	result := responses.HubtelPaymentRequestApiResponseData{}
	resp := responses.HubtelPaymentRequestResponseDTO{
		Success:    responseCode,
		StatusDesc: responseMessage,
		Result:     &result,
	}

	requestPaymentResp, err := HubtelRequestViaMobileMoney(c, momoRequest)

	if err != nil {
		logs.Error("Error making payment request via mobile money: ", err)
		return responses.HubtelPaymentRequestResponseDTO{
			Success:    false,
			StatusDesc: "Error processing request",
			Result:     nil,
		}, err
	}

	if requestPaymentResp.ResponseCode == "0001" {
		if status, err := models.GetStatusByName("PENDING"); err == nil {
			req.Payment.Status = status

			if err := models.UpdatePaymentsById(&req.Payment); err == nil {
				logs.Info("Payment status updated to PENDING for Payment ID: ", req.Payment.PaymentId)
				responseCode = true
				responseMessage = "Payment request initiated successfully"
			} else {
				logs.Error("Error updating payment status to PENDING for Payment ID: ", req.Payment.PaymentId, " Error: ", err)
				responseCode = false
				responseMessage = "Error updating payment status"
			}
		}
	}

	resp = responses.HubtelPaymentRequestResponseDTO{
		Success:    responseCode,
		StatusDesc: responseMessage,
		Result:     &result,
	}
	return resp, nil
}

func NameInquiryViaMobileMoney(c *beego.Controller, req requests.HubtelNameInquiryRequestDTO) (responses.HubtelNameInquiryResponseDTO, error) {
	//do something

	responseCode := false
	responseMessage := "Error processing request"
	result := responses.HubtelNameInquiryApiResponseData{}
	resp := responses.HubtelNameInquiryResponseDTO{
		Success:    responseCode,
		StatusDesc: responseMessage,
		Result:     &result,
	}

	requestNameInquiryResp, err := HubtelNameInquiry(c, req.CustomerMsisdn, req.Channel)

	if err != nil {
		logs.Error("Error making name inquiry via mobile money: ", err)
		return responses.HubtelNameInquiryResponseDTO{
			Success:    false,
			StatusDesc: "Error processing request",
			Result:     nil,
		}, err
	}

	if requestNameInquiryResp.ResponseCode == "0000" {
		responseCode = true
		responseMessage = "Name inquiry successful"
		result = responses.HubtelNameInquiryApiResponseData{
			Display: requestNameInquiryResp.Data[0].Display,
			Value:   requestNameInquiryResp.Data[0].Value,
			Amount:  requestNameInquiryResp.Data[0].Amount,
		}
	}

	resp = responses.HubtelNameInquiryResponseDTO{
		Success:    responseCode,
		StatusDesc: responseMessage,
		Result:     &result,
	}
	return resp, nil
}
