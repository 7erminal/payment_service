package controllers

import (
	"encoding/json"
	"payment_service/controllers/functions"
	"payment_service/models"
	"payment_service/structs/requests"
	"payment_service/structs/responses"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Request_moneyController operations for Request_money
type Request_moneyController struct {
	beego.Controller
}

// URLMapping ...
func (c *Request_moneyController) URLMapping() {
	c.Mapping("RequestMoneyViaMomo", c.RequestMoneyViaMomo)
}

// RequestMoneyViaMomo ...
// @Title RequestMoneyViaMomo
// @Description create Request_money
// @Param	body		body 	requests.MomoPaymentRequestDTO	true		"body for Request_money content"
// @Success 201 {object} models.Request_money
// @Failure 403 body is empty
// @router /momo/ [post]
func (c *Request_moneyController) RequestMoneyViaMomo() {
	var v requests.MomoPaymentRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	responseCode := 401
	responseMessage := "Error processing request"
	resp := responses.RequestMoneyResponseDTO{
		StatusCode: responseCode,
		Result:     nil,
		StatusDesc: responseMessage,
	}
	paymentId := v.PaymentId
	logs.Info("Processing payment request for Payment ID: ", paymentId)

	id, perr := strconv.ParseInt(paymentId, 10, 64)
	if perr != nil {
		logs.Error("Invalid payment ID: ", perr)
		resp = responses.RequestMoneyResponseDTO{StatusCode: 400, Result: nil, StatusDesc: "Invalid payment ID: " + perr.Error()}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	if payment, err := models.GetPaymentsById(id); err == nil {
		logs.Info("Payment found: ", payment)
		if network, err := models.GetNetworksByCode(v.Channel); err == nil {
			if v.Operator == "HUBTEL" {
				customerName := v.CustomerName
				callbackurl := ""
				if cbr, err := models.GetApplication_propertyByCode("HUBTEL_PAYMENT_CALLBACK_URL"); err == nil {
					callbackurl = cbr.PropertyValue
				} else {
					logs.Error("Failed to get callback URL: %v", err)
				}

				if status, err := models.GetStatusByName("PENDING"); err == nil {
					var payment_history models.Payment_history = models.Payment_history{
						Payment:      payment,
						Status:       status.StatusId,
						Service:      "MOBILEMONEY",
						Narration:    "Requesting money via Mobile Money for Payment ID " + paymentId,
						Reference:    v.Channel,
						DateCreated:  time.Now(),
						DateModified: time.Now(),
						CreatedBy:    1,
						ModifiedBy:   1,
						Active:       1,
					}
					if _, err := models.AddPayment_history(&payment_history); err == nil {
						momoRequest := requests.MomoPaymentApiRequestDTO{
							CustomerName:       customerName,
							CustomerMsisdn:     v.CustomerMsisdn,
							CustomerEmail:      v.CustomerEmail,
							Channel:            network.NetworkReferenceId,
							Amount:             float32(v.Amount),
							PrimaryCallbackUrl: callbackurl,
							Description:        "Payment for " + customerName,
							ClientReference:    v.ClientReference,
						}

						payment := responses.RequestMoneyDataResponse{
							PaymentId:          v.PaymentId,
							Amount:             payment.Amount,
							SenderAccount:      payment.SenderAccount,
							ReceiverAccount:    payment.ReceiverAccount,
							ReferenceNumber:    payment.ReferenceNumber,
							Description:        payment.Narration,
							AmountAfterCharges: payment.OtherCharge,
							AmountCharged:      payment.Charge,
							PaymentDate:        payment.DateCreated.Format("2006-01-02 15:04:05"),
						}

						if hubtelResp, err := functions.PaymentRequestViaMobileMoney(&c.Controller, momoRequest); err == nil {
							logs.Info("Hubtel payment request response: ", hubtelResp)
							if hubtelResp.Success {
								responseCode = 200
								responseMessage = "Payment request successful"
								payment.ReferenceNumber = hubtelResp.Result.ClientReference
								payment.Description = hubtelResp.Result.Description
								payment.AmountAfterCharges = hubtelResp.Result.AmountAfterCharges
								payment.AmountCharged = hubtelResp.Result.AmountCharged
								payment.PaymentDate = hubtelResp.Result.PaymentDate
								resp = responses.RequestMoneyResponseDTO{StatusCode: responseCode, Result: &payment, StatusDesc: responseMessage}
							} else {
								responseMessage = "Payment request failed! " + hubtelResp.StatusDesc
								resp = responses.RequestMoneyResponseDTO{StatusCode: responseCode, Result: &payment, StatusDesc: responseMessage}
							}
						}

					} else {
						logs.Error("Failed to create payment record: %v", err)
						resp = responses.RequestMoneyResponseDTO{StatusCode: 807, Result: nil, StatusDesc: "Order error! " + err.Error()}
					}
				} else {
					logs.Error("Unable to get status PENDING: %v", err)
					resp = responses.RequestMoneyResponseDTO{StatusCode: 808, Result: nil, StatusDesc: "Order error! " + err.Error()}
				}
			}

		} else {
			logs.Error("Unable to get network ", err.Error())
			resp = responses.RequestMoneyResponseDTO{StatusCode: 806, Result: nil, StatusDesc: "Order error! " + err.Error()}
		}
	} else {
		logs.Error("Unable to find payment ", err.Error())
		resp = responses.RequestMoneyResponseDTO{StatusCode: 805, Result: nil, StatusDesc: "Order error! " + err.Error()}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
