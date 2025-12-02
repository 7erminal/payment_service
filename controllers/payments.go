package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
	"payment_service/controllers/functions"
	"payment_service/models"
	"payment_service/structs/requests"
	"payment_service/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// PaymentsController operations for Payments
type PaymentsController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaymentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UploadPaymentProof", c.UploadPaymentProof)
	c.Mapping("GetPaymentCount", c.GetPaymentCount)
	c.Mapping("NameInquiry", c.NameInquiry)
	c.Mapping("SendMoneyViaMomo", c.SendMoneyViaMomo)
}

// Post ...
// @Title Post
// @Description create Payments
// @Param	body		body 	requests.PaymentRequest2DTO	true		"body for Payments content"
// @Success 201 {int} requests.PaymentResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *PaymentsController) Post() {
	var v requests.PaymentRequest2DTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Request received ", v)
	logs.Info("Transaction ID is ", v.TransactionId)

	reqText, err := json.Marshal(v)
	if err != nil {
		c.Data["json"] = "Invalid request format"
		c.ServeJSON()
		return
	}

	statusCode := "PENDING"

	serviceCode := "PAYMENT"
	if service, err := models.GetServicesByCode(serviceCode); err == nil {
		logs.Info("Service fetched successfully")
		status, err := models.GetStatus_codesByCode(statusCode)
		if err == nil {
			var sender models.Customers
			if s, err := models.GetCustomerById(v.Sender); err == nil {
				sender = *s
			} else {
				logs.Error("Error getting customer ", err.Error())
			}

			req := models.Request{
				CustId:          &sender,
				Request:         string(reqText),
				RequestType:     service.ServiceName,
				RequestStatus:   status.StatusDescription,
				RequestAmount:   float64(v.Amount),
				RequestResponse: "",
				RequestDate:     time.Now(),
				DateCreated:     time.Now(),
				DateModified:    time.Now(),
			}
			if _, err := models.AddRequest(&req); err == nil {
				if paymentMethod, err := models.GetPayment_methodsByName(v.PaymentMethod); err == nil {

					var receiver models.Users
					if u, err := models.GetUsersById(v.Reciever); err == nil {
						receiver = *u
					} else {
						logs.Error("Error getting user ", err.Error())
					}
					if status, err := models.GetStatusByName("PENDING"); err == nil {
						logs.Info("Payment reference number is " + v.ReferenceNumber)
						var transaction *models.Transactions
						if v.TransactionId != "" {
							if tid, err := strconv.ParseInt(v.TransactionId, 10, 64); err == nil {
								transaction = &models.Transactions{TransactionId: tid}
							} else {
								logs.Error("Invalid TransactionId: ", err)
							}
						}

						var payment models.Payments = models.Payments{
							Transaction:     transaction,
							PaymentProof:    v.PaymentProofUrl,
							ReferenceNumber: v.ReferenceNumber,
							Request:         &req,
							InitiatedBy:     v.InitiatedBy,
							Sender:          &sender,
							Reciever:        &receiver,
							Amount:          float64(v.Amount),
							PaymentMethod:   paymentMethod,
							PaymentCurrency: v.Currency,
							Status:          status,
							PaymentAccount:  "",
							DateCreated:     time.Now(),
							DateModified:    time.Now(),
							DateProcessed:   time.Now(),
							CreatedBy:       v.InitiatedBy,
							ModifiedBy:      v.InitiatedBy,
							Active:          1,
							Service:         v.Service,
							ServiceNetwork:  v.ServiceNetwork,
							ServicePackage:  v.ServicePackage,
							SenderAccount:   v.SenderAccount,
							ReceiverAccount: v.ReceiverAccount,
						}
						if _, err := models.AddPayments(&payment); err == nil {
							// Send to Account service to debit and credit
							logs.Info("Payment added successfully")
							logs.Info(payment)

							var payment_history models.Payment_history = models.Payment_history{
								Payment:      &payment,
								Status:       payment.Status.StatusId,
								Service:      payment.Service,
								Narration:    "Payment initiated",
								Reference:    v.Network,
								DateCreated:  time.Now(),
								DateModified: time.Now(),
								CreatedBy:    v.InitiatedBy,
								ModifiedBy:   v.InitiatedBy,
								Active:       1,
							}
							if _, err := models.AddPayment_history(&payment_history); err == nil {
								logs.Info("Payment history added successfully")
								logs.Info("Checking if call back is required...")
								callbackUrl := ""
								if v.CallThirdParty {
									logs.Info("Callback required")
									operatorCaps := strings.ToUpper(v.Operator)
									serviceCaps := strings.ToUpper("PAYMENT")
									operator, err := models.GetOperatorByName(operatorCaps)
									if err == nil {
										appProperty, err := models.GetApplication_propertyByCode(strings.ToUpper(operator.OperatorName) + "_" + serviceCaps + "_CALLBACK_URL")

										if err == nil && appProperty != nil {
											callbackUrl = appProperty.PropertyValue
											logs.Info("Callback URL found: " + callbackUrl)
										} else {
											logs.Error("Callback URL not found for operator "+operator.OperatorName+" and service "+serviceCaps+": ", err)
										}
									} else {
										logs.Error("Operator not found for callback: ", err)
									}
								}
								paymentResp := responses.PaymentResponse{
									PaymentId:       strconv.FormatInt(payment.PaymentId, 10),
									Sender:          payment.Sender.FullName,
									Reciever:        payment.Reciever.FullName,
									Amount:          payment.Amount,
									Commission:      payment.Commission,
									Charge:          payment.Charge,
									OtherCharge:     payment.OtherCharge,
									PaymentAmount:   payment.PaymentAmount,
									PaymentMethod:   payment.PaymentMethod,
									PaymentProof:    payment.PaymentProof,
									Status:          payment.Status,
									Service:         payment.Service,
									SenderAccount:   payment.SenderAccount,
									ReceiverAccount: payment.ReceiverAccount,
									ReferenceNumber: payment.ReferenceNumber,
									DateCreated:     payment.DateCreated,
									DateModified:    payment.DateModified,
									ProcessedDate:   payment.DateProcessed,
									Active:          payment.Active,
									CallbackUrl:     callbackUrl,
								}
								var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 200, Payment: &paymentResp, StatusDesc: "Payment successfully initiated!"}

								c.Ctx.Output.SetStatus(200)
								c.Data["json"] = resp
							} else {
								logs.Error("Unable to add payment history ", err.Error())
								var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
								c.Data["json"] = resp
							}
						} else {
							logs.Error("Unable to add payment ", err.Error())
							var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
							c.Data["json"] = resp
						}
					}
				} else {
					logs.Error("Unable to get payment method ", err.Error())
					var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
					c.Data["json"] = resp
				}
			} else {
				logs.Error("Unable to add request ", err.Error())
				var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Unable to get status ", err.Error())
			var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Unable to get service ", err.Error())
		var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// SendMoney ...
// @Title Send Money via Mobile Money
// @Description create Send Money via Mobile Money
// @Param	body		body 	requests.MomoPaymentRequestDTO	true		"body for Request_money content"
// @Success 201 {object} models.Request_money
// @Failure 403 body is empty
// @router /momo/ [post]
func (c *PaymentsController) SendMoneyViaMomo() {
	var v requests.MomoPaymentRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	responseCode := 401
	responseMessage := "Error processing request"
	resp := responses.SendMoneyResponseDTO{
		StatusCode: responseCode,
		Result:     nil,
		StatusDesc: responseMessage,
	}
	paymentId := v.PaymentId
	logs.Info("Processing payment request for Payment ID: ", paymentId)

	id, perr := strconv.ParseInt(paymentId, 10, 64)
	if perr != nil {
		logs.Error("Invalid payment ID: ", perr)
		resp = responses.SendMoneyResponseDTO{StatusCode: 400, Result: nil, StatusDesc: "Invalid payment ID: " + perr.Error()}
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
						Narration:    "Sending money via Mobile Money for Payment ID " + paymentId,
						Reference:    v.Channel,
						DateCreated:  time.Now(),
						DateModified: time.Now(),
						CreatedBy:    1,
						ModifiedBy:   1,
						Active:       1,
					}
					if _, err := models.AddPayment_history(&payment_history); err == nil {
						momoRequest := requests.MomoPaymentApiRequestDTO{
							Payment:            *payment,
							CustomerName:       customerName,
							CustomerMsisdn:     v.CustomerMsisdn,
							CustomerEmail:      v.CustomerEmail,
							Channel:            network.NetworkReferenceId,
							Amount:             float32(v.Amount),
							PrimaryCallbackUrl: callbackurl,
							Description:        "Payment for " + customerName,
							ClientReference:    v.ClientReference,
						}

						payment := responses.SendMoneyDataResponse{
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

						if hubtelResp, err := functions.PaymentSendViaMobileMoney(&c.Controller, momoRequest); err == nil {
							logs.Info("Hubtel send payment request response: ", hubtelResp)
							if hubtelResp.Success {
								responseCode = 200
								responseMessage = "Payment request successful"
								payment.ReferenceNumber = hubtelResp.Result.ClientReference
								payment.Description = hubtelResp.Result.Description
								payment.AmountAfterCharges = hubtelResp.Result.AmountAfterCharges
								payment.AmountCharged = hubtelResp.Result.AmountCharged
								resp = responses.SendMoneyResponseDTO{StatusCode: responseCode, Result: &payment, StatusDesc: responseMessage}
							} else {
								responseMessage = "Payment request failed! " + hubtelResp.StatusDesc
								resp = responses.SendMoneyResponseDTO{StatusCode: responseCode, Result: &payment, StatusDesc: responseMessage}
							}
						}

					} else {
						logs.Error("Failed to create payment record: %v", err)
						resp = responses.SendMoneyResponseDTO{StatusCode: 807, Result: nil, StatusDesc: "Order error! " + err.Error()}
					}
				} else {
					logs.Error("Unable to get status PENDING: %v", err)
					resp = responses.SendMoneyResponseDTO{StatusCode: 808, Result: nil, StatusDesc: "Order error! " + err.Error()}
				}
			}

		} else {
			logs.Error("Unable to get network ", err.Error())
			resp = responses.SendMoneyResponseDTO{StatusCode: 806, Result: nil, StatusDesc: "Order error! " + err.Error()}
		}
	} else {
		logs.Error("Unable to find payment ", err.Error())
		resp = responses.SendMoneyResponseDTO{StatusCode: 805, Result: nil, StatusDesc: "Order error! " + err.Error()}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// Name Inquiry ...
// @Title Name Inquiry
// @Description Get the name associated with a mobile money number
// @Param	body		body 	requests.NameInquiryRequestDTO	true		"body for Name Inquiry content"
// @Success 200 {object} responses.NameInquiryResponseDTO
// @Failure 403 body is empty
// @router /name-inquiry [post]
func (c *PaymentsController) NameInquiry() {
	var v requests.NameInquiryRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Request received ", v)
	logs.Info("Network is ", v.Channel)

	if network, err := models.GetNetworksByCode(v.Channel); err == nil {
		hubtelRequest := requests.HubtelNameInquiryRequestDTO{
			CustomerMsisdn: v.CustomerMsisdn,
			Channel:        network.NetworkCode,
		}

		if hubtelResp, err := functions.NameInquiryViaMobileMoney(&c.Controller, hubtelRequest); err == nil {
			logs.Info("Hubtel name inquiry response: ", hubtelResp)
			if hubtelResp.Success {
				resp := responses.HubtelNameInquiryResponseDTO{Success: true, Result: hubtelResp.Result, StatusDesc: "Name inquiry successful!"}
				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = resp
			} else {
				resp := responses.HubtelNameInquiryResponseDTO{Success: false, Result: nil, StatusDesc: "Name inquiry failed! " + hubtelResp.StatusDesc}
				c.Ctx.Output.SetStatus(805)
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error making name inquiry via mobile money: ", err)
			resp := responses.HubtelNameInquiryResponseDTO{Success: false, Result: nil, StatusDesc: "Error processing request"}
			c.Ctx.Output.SetStatus(806)
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Unable to get network ", err.Error())
		resp := responses.HubtelNameInquiryResponseDTO{Success: false, Result: nil, StatusDesc: "Order error! " + err.Error()}
		c.Ctx.Output.SetStatus(806)
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Upload Payment Proof ...
// @Title Upload Payment Proof
// @Description Upload a payment proof
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {int} responses.StringResponseFDTO
// @Failure 403 body is empty
// @router /upload-payment-proof [post]
func (c *PaymentsController) UploadPaymentProof() {
	// var v models.Item_images
	file, header, err := c.GetFile("Image")
	logs.Info("Data received is ", file)

	contentType := c.Ctx.Input.Header("Content-Type")
	logs.Info("Content-Type of incoming request:", contentType)

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Error("Failed to get the file ", err)
		c.ServeJSON()
		return
	}
	defer file.Close()

	// Check the file size
	fileInfo := header.Size
	logs.Info("Received file size:", fileInfo)

	// Save the uploaded file
	fileName := filepath.Base(header.Filename)
	logs.Info("File Name Extracted is ", fileName, "Time now is ", time.Now().Format("20060102150405"))
	filePath := "/uploads/payments/" + time.Now().Format("20060102150405") + fileName // Define your file path
	logs.Info("File Path Extracted is ", filePath)
	err = c.SaveToFile("Image", "../images"+filePath)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.Error("Error saving file", err)
		// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
		errorMessage := "Error: Failed to save the image file"

		resp := responses.StringResponseDTO{StatusCode: http.StatusInternalServerError, Value: errorMessage, StatusDesc: "Internal Server Error"}

		c.Data["json"] = resp
		c.ServeJSON()
		return
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: filePath, StatusDesc: "Images uploaded successfully"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Payments by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Payments
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PaymentsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPaymentsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Payments
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Payments
// @Failure 403
// @router / [get]
func (c *PaymentsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var search = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
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

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	message := "An error occurred adding this audit request"
	statusCode := 308

	l, err := models.GetAllPayments(query, fields, sortby, order, offset, limit, search)
	if err != nil {
		logs.Info("Error fetching expenses ", err.Error())
		message = "Error fetching expenses."
		statusCode = 608
		resp := responses.PaymentsResponseDTO{StatusCode: statusCode, Payments: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		if l == nil {
			l = []interface{}{}
		}
		statusCode = 200
		message = "Expenses fetched successfully"
		resp := responses.PaymentsResponseDTO{StatusCode: statusCode, Payments: &l, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Payments
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Payments	true		"body for Payments content"
// @Success 200 {object} models.Payments
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PaymentsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Payments{PaymentId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePaymentsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Payments
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PaymentsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeletePayments(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetPaymentCount ...
// @Title Get payment Count
// @Description get Count of payments
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/ [get]
func (c *PaymentsController) GetPaymentCount() {
	// q, err := models.GetItemsById(id)
	var query = make(map[string]string)
	var search = make(map[string]string)

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
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

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	v, err := models.GetPaymentCount(query, search)
	count := strconv.FormatInt(v, 10)
	if err != nil {
		logs.Error("Error fetching count of expenses ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}
