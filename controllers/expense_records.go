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

// Expense_recordsController operations for Expense_records
type Expense_recordsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Expense_recordsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Expense_records
// @Param	body		body 	requests.ExpenseRequest	true		"body for Expense_records content"
// @Success 201 {int} models.Expense_records
// @Failure 403 body is empty
// @router / [post]
func (c *Expense_recordsController) Post() {
	var v requests.ExpenseRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var expenseDate time.Time
	proceed := false
	message := "An error occurred adding this audit request"
	statusCode := 308

	var allowedDateList [4]string = [4]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000"}

	for _, date_ := range allowedDateList {
		logs.Debug("About to convert ", v.Date)
		logs.Debug("About to convert ", c.Ctx.Input.Query("Dob"))
		// Convert dob string to date
		tDate, error := time.Parse(date_, v.Date)

		if error != nil {
			logs.Error("Error parsing date", error)
			message = "Unable to determine date changed"
			proceed = false
		} else {
			logs.Error("Date converted to time successfully", tDate)
			expenseDate = tDate
			proceed = true

			break
		}
	}

	if proceed {
		currency := models.Currencies{}

		if v.Currency != 0 {
			logs.Info("Currency is not 0")
			if cur, err := models.GetCurrenciesById(v.Currency); err == nil {
				currency = *cur
			}
		}

		if category, err := models.GetPayment_categoriesById(v.Category); err == nil {
			if paymentMethod, err := models.GetPayment_methodsById(v.PaymentMethod); err == nil {
				expense := models.Expense_records{ExpenseDate: expenseDate, Category: category, Description: v.Description, Amount: v.Amount, Currency: &currency, PaymentMethod: paymentMethod, Active: 1}

				if _, err := models.AddExpense_records(&expense); err == nil {
					statusCode = 200
					message = "Expense updated successfully"
					resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: &expense, StatusDesc: message}
					c.Data["json"] = resp
				} else {
					logs.Info("Error adding expense ", err.Error())
					message = "Error adding expense"
					statusCode = 608
					resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: nil, StatusDesc: message}
					c.Data["json"] = resp
				}
			} else {
				logs.Info("Error adding expense ", err.Error())
				message = "Error adding expense"
				statusCode = 608
				resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: nil, StatusDesc: message}
				c.Data["json"] = resp
			}
		} else {
			logs.Info("Error adding expense ", err.Error())
			message = "Error adding expense. Category not found"
			statusCode = 608
			resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: nil, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		message = "Error adding expense. Date provided is not valid"
		statusCode = 608
		resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Expense_records by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Expense_records
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Expense_recordsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetExpense_recordsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Expense_records
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Expense_records
// @Failure 403
// @router / [get]
func (c *Expense_recordsController) GetAll() {
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

	l, err := models.GetAllExpense_records(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Expense_records
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Expense_records	true		"body for Expense_records content"
// @Success 200 {object} models.Expense_records
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Expense_recordsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Expense_records{ExpenseRecordId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateExpense_recordsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Expense_records
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Expense_recordsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteExpense_records(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
