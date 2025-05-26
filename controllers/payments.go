package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
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
}

// Post ...
// @Title Post
// @Description create Payments
// @Param	body		body 	requests.PaymentRequestDTO	true		"body for Payments content"
// @Success 201 {int} requests.PaymentRequestDTO
// @Failure 403 body is empty
// @router / [post]
func (c *PaymentsController) Post() {
	var v requests.PaymentRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Request received ", v)
	logs.Info("Transaction ID is ", v.TransactionId)

	if transaction, err := models.GetTransactionsById(v.TransactionId); err == nil {
		if paymentMethod, err := models.GetPayment_methodsById(v.PaymentMethod); err == nil {
			var sender models.Customers
			if s, err := models.GetCustomerById(v.Sender); err == nil {
				sender = *s
			} else {
				logs.Error("Error getting customer ", err.Error())
			}

			var receiver models.Users
			if u, err := models.GetUsersById(v.Reciever); err == nil {
				receiver = *u
			} else {
				logs.Error("Error getting user ", err.Error())
			}
			if status, err := models.GetStatusByName("PENDING"); err == nil {
				var payment models.Payments = models.Payments{Transaction: transaction, PaymentProof: v.PaymentProofUrl, ReferenceNumber: v.ReferenceNumber, InitiatedBy: v.InitiatedBy, Sender: &sender, Reciever: &receiver, Amount: float64(v.Amount), PaymentMethod: paymentMethod, Status: status, PaymentAccount: 0, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: v.InitiatedBy, ModifiedBy: v.InitiatedBy, Active: 1}
				if _, err := models.AddPayments(&payment); err == nil {
					// Send to Account service to debit and credit
					logs.Info("Payment added successfully")
					logs.Info(payment)

					var payment_history models.Payment_history = models.Payment_history{PaymentId: payment.PaymentId, Status: payment.Status.StatusId, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: v.InitiatedBy, ModifiedBy: v.InitiatedBy, Active: 1}
					if _, err := models.AddPayment_history(&payment_history); err == nil {
						var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 200, Payment: &payment, StatusDesc: "Payment successfully initiated!"}
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
		logs.Error("Unable to get transaction ", err.Error())
		var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
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
