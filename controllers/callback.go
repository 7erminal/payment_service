package controllers

import (
	"encoding/json"
	"payment_service/models"
	"payment_service/structs/requests"
	"payment_service/structs/responses"
	"strconv"
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

	responseCode := false
	responseMessage := "Invalid request"

	// Handle successful callback
	transactionId := ""
	if v.ClientReference != nil {
		logs.Info("Transaction ID found in request: ", *v.ClientReference)
		transactionId = *v.ClientReference
	}
	logs.Info("About to get transaction by ID: ", transactionId)
	id, err := strconv.ParseInt(transactionId, 10, 64)
	if err != nil {
		logs.Error("Invalid transaction ID: %v", err)
		responseCode = false
		responseMessage = "Invalid transaction ID"
		resp := responses.CallbackResponse{
			StatusCode:    responseCode,
			StatusMessage: responseMessage,
			Result:        nil,
		}
		c.Data["json"] = resp
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	if resp, err := models.GetPaymentsById(id); err == nil {
		logs.Info("Request ID: ", resp.PaymentId)
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
			} else {
				c.Data["json"] = map[string]string{"error": "Status code not found"}
				c.Ctx.Output.SetStatus(404)
			}

			if err := models.UpdatePaymentsById(resp); err != nil {
				logs.Info("Failed to update transaction status: %v", err)
				responseCode = false
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

				responseCode = true
				responseMessage = "Transaction updated successfully"
				payment := responses.PaymentResponse{
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
					DateCreated:     resp.DateCreated,
					DateModified:    resp.DateModified,
					ProcessedDate:   resp.DateProcessed,
					Active:          resp.Active,
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
			logs.Info("Transaction not found for ID: %s", transactionId)
			responseCode = false
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
		responseCode = false
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
