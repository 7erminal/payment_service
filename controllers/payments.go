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
	logs.Info("Request received")

	var payment models.Payments = models.Payments{TransactionId: v.TransactionId, InitiatedBy: v.InitiatedBy, Sender: v.Sender, Reciever: v.Reciever, Amount: float64(v.Amount), PaymentMethod: v.PaymentMethod, Status: 2, PaymentAccount: 0, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: v.InitiatedBy, ModifiedBy: v.InitiatedBy, Active: 1}
	if _, err := models.AddPayments(&payment); err == nil {
		// Send to Account service to debit and credit
		logs.Info("Payment added successfully")
		logs.Info(payment)

		var payment_history models.Payment_history = models.Payment_history{PaymentId: payment.PaymentId, Status: payment.Status, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: v.InitiatedBy, ModifiedBy: v.InitiatedBy, Active: 1}
		if _, err := models.AddPayment_history(&payment_history); err == nil {
			var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 200, Payment: &payment, StatusDesc: "Payment successfully initiated!"}
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = resp
		} else {
			var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.PaymentResponseDTO = responses.PaymentResponseDTO{StatusCode: 806, Payment: nil, StatusDesc: "Order error! " + err.Error()}
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

	l, err := models.GetAllPayments(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
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
