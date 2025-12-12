package controllers

import (
	"encoding/json"
	"errors"
	"payment_service/models"
	"payment_service/structs/requests"
	"payment_service/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// CallbackController operations for Callback
type CallbackController struct {
	beego.Controller
}

// URLMapping ...
func (c *CallbackController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Callback
// @Param	body		body 	requests.CallbackRequest	true		"body for Callback content"
// @Success 201 {object} responses.CallbackResponse
// @Failure 403 body is empty
// @router /process [post]
func (c *CallbackController) Post() {
	var v requests.CallbackRequest
	logs.Info("Received callback request: ", string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	responseCode := 606
	responseMessage := "Invalid request"

	// Handle successful callback
	// transactionId := ""
	// if v.ClientReference != nil {
	// 	logs.Info("Transaction ID found in request: ", *v.ClientReference)
	// 	transactionId = *v.ClientReference
	// }
	// logs.Info("About to get transaction by ID: ", transactionId)
	// id, err := strconv.ParseInt(transactionId, 10, 64)
	// if err != nil {
	// 	logs.Error("Invalid transaction ID: %v", err)
	// 	responseCode = false
	// 	responseMessage = "Invalid transaction ID"
	// 	resp := responses.CallbackResponse{
	// 		StatusCode:    responseCode,
	// 		StatusMessage: responseMessage,
	// 		Result:        nil,
	// 	}
	// 	c.Data["json"] = resp
	// 	c.Ctx.Output.SetStatus(400)
	// 	c.ServeJSON()
	// 	return
	// }
	if resp, err := models.GetPaymentsByTxnReference(*v.ClientReference); err == nil {
		logs.Info("Request ID: ", &resp.PaymentId)
		if resp != nil {
			// Update the transaction status
			statusCode := v.Status
			logs.Info("Updating transaction status to: ", statusCode)

			status, err := models.GetStatusByName(statusCode)
			if err == nil {
				resp.Status = status
				resp.DateModified = time.Now()
				if v.ExternalTransactionId != nil {
					resp.ReferenceNumber = *v.ExternalTransactionId
				}
				resp.DateProcessed = time.Now()
				resp.Charge = v.Charges
				resp.OtherCharge = v.AmountCharged
				resp.PaymentAmount = v.Amount
				resp.Commission, _ = strconv.ParseFloat(v.Commission, 64)
			} else {
				c.Data["json"] = map[string]string{"error": "Status code not found"}
				c.Ctx.Output.SetStatus(404)
			}

			if err := models.UpdatePaymentsById(resp); err != nil {
				logs.Info("Failed to update transaction status: %v", err)
				responseCode = 607
				responseMessage = "Failed to update transaction status"
				resp := responses.CallbackResponse{
					StatusCode:    responseCode,
					StatusMessage: responseMessage,
					Result:        nil,
				}
				c.Data["json"] = resp
				c.Ctx.Output.SetStatus(200)
			} else {
				// c.Data["json"] = map[string]string{"message": "Transaction updated successfully"}

				// Update payment history
				var fields []string
				var sortby []string
				var order []string
				var query = make(map[string]string)
				var limit int64 = 10
				var offset int64

				querySearch := "Payment__PaymentId:" + strconv.FormatInt(resp.PaymentId, 10)
				// query: k:v,k:v
				if v := querySearch; v != "" {
					for _, cond := range strings.Split(v, ",") {
						kv := strings.SplitN(cond, ":", 2)
						if len(kv) != 2 {
							c.Data["json"] = errors.New("Error: invalid query key/value pair")
							c.ServeJSON()
							return
						}
						k, v := kv[0], kv[1]
						query[k] = v
					}
				}

				var paymentHistoryResponses []responses.PaymentHistoryResponse

				paymentHist, err := models.GetAllPayment_history(query, fields, sortby, order, offset, limit)
				if err == nil {
					for _, v := range paymentHist {
						logs.Info("Payment history found: ", v)
						paymentHistObj := v.(models.Payment_history)

						paymentHistObj.Status = status.StatusId
						paymentHistObj.DateModified = time.Now()

						if err := models.UpdatePayment_historyById(&paymentHistObj); err != nil {
							logs.Error("Failed to update payment history: %v", err)
						} else {
							logs.Info("Payment history updated successfully for Payment ID: ", resp.PaymentId)
						}

						paymentHistoryResponse := responses.PaymentHistoryResponse{
							PaymentHistoryId: paymentHistObj.PaymentHistoryId,
							PaymentId:        paymentHistObj.Payment.PaymentId,
							Status:           status.Status,
							Service:          paymentHistObj.Service,
							Narration:        paymentHistObj.Narration,
							Reference:        paymentHistObj.Reference,
							DateCreated:      paymentHistObj.DateCreated,
							DateModified:     paymentHistObj.DateModified,
							CreatedBy:        paymentHistObj.CreatedBy,
							ModifiedBy:       paymentHistObj.ModifiedBy,
							Active:           paymentHistObj.Active,
						}
						paymentHistoryResponses = append(paymentHistoryResponses, paymentHistoryResponse)
					}

				} else {
					logs.Error("Failed to retrieve payment history: %v", err)
				}

				// Update request with callback data
				resText, err := json.Marshal(v)
				if err != nil {
					logs.Error("Failed to marshal callback request: %v", err)
					// c.Data["json"] = "Invalid request format"
					// c.ServeJSON()
					// return
				}

				logs.Info("Callback response text: %s", string(resText))
				logs.Info("Updating request", resp.Request.RequestId, " with callback response")
				if request, err := models.GetRequestById(resp.Request.RequestId); err == nil {
					logs.Info("Found request: ", request.RequestId)
					request.CallbackResponse = string(resText)

					request.DateModified = time.Now()
					if err := models.UpdateRequestById(request); err != nil {
						logs.Error("Failed to update request: %v", err)
						// c.Data["json"] = "Failed to update request"
						// c.ServeJSON()
						// return
					} else {
						logs.Info("Request updated successfully with callback response")
					}
				} else {
					logs.Error("Failed to retrieve request by ID: %v", err)
				}

				logs.Info("Returniing transaction ID: %s", resp.TransactionId)
				respJSON, err := json.Marshal(resp)
				if err != nil {
					logs.Error("Failed to marshal response: %v", err)
				} else {
					logs.Info("Response: %s", string(respJSON))
				}
				responseCode = 200
				responseMessage = "Transaction updated successfully"
				payment := responses.PaymentResponse{
					TransactionId:   resp.TransactionId,
					Sender:          resp.Sender.FullName,
					Reciever:        resp.Reciever.FullName,
					Amount:          resp.Amount,
					Commission:      resp.Commission,
					Charge:          resp.Charge,
					OtherCharge:     resp.OtherCharge,
					PaymentAmount:   resp.PaymentAmount,
					PaymentMethod:   resp.PaymentMethod,
					PaymentProof:    resp.PaymentProof,
					Status:          resp.Status,
					Service:         resp.Service,
					SenderAccount:   resp.SenderAccount,
					ReceiverAccount: resp.ReceiverAccount,
					ReferenceNumber: resp.ReferenceNumber,
					ServiceNetwork:  resp.ServiceNetwork,
					ServicePackage:  resp.ServicePackage,
					DateCreated:     resp.DateCreated,
					DateModified:    resp.DateModified,
					ProcessedDate:   resp.DateProcessed,
					Active:          resp.Active,
					PaymentHistory:  &paymentHistoryResponses,
				}
				resp := responses.CallbackResponse{
					StatusCode:    responseCode,
					StatusMessage: responseMessage,
					Result:        &payment,
				}
				c.Data["json"] = resp
				c.Ctx.Output.SetStatus(200)
			}
		} else {
			logs.Info("Transaction not found for ID: %s", v.TransactionId)
			responseCode = 608
			responseMessage = "Transaction not found"
			resp := responses.CallbackResponse{
				StatusCode:    responseCode,
				StatusMessage: responseMessage,
				Result:        nil,
			}
			c.Data["json"] = resp
			// c.Data["json"] = map[string]string{"error": "Transaction not found"}
			c.Ctx.Output.SetStatus(200)
		}
	} else {
		c.Data["json"] = map[string]string{"error": "Failed to retrieve transaction"}
		logs.Info("Failed to retrieve transaction: %s", err.Error())
		responseCode = 609
		responseMessage = "Failed to retrieve transaction"
		resp := responses.CallbackResponse{
			StatusCode:    responseCode,
			StatusMessage: responseMessage,
			Result:        nil,
		}
		c.Data["json"] = resp
		// c.Data["json"] = map[string]string{"error": "Transaction not found"}
		c.Ctx.Output.SetStatus(200)
	}

	c.ServeJSON()
}
