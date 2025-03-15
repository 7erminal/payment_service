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
	c.Mapping("GetExpenseRecordCount", c.GetExpenseRecordCount)
	c.Mapping("GetAllByBranch", c.GetAllByBranch)
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

	rexpenseDate := c.Ctx.Input.Query("Date")
	rdescription := c.Ctx.Input.Query("Description")
	rcurrency := c.Ctx.Input.Query("Currency")
	rpaymentmethod := c.Ctx.Input.Query("PaymentMethod")
	ramount := c.Ctx.Input.Query("Amount")
	ruser := c.Ctx.Input.Query("AddedBy")
	rcategory := c.Ctx.Input.Query("Category")

	// image of user received
	file, header, err := c.GetFile("ReceiptImage")
	var filePath string = ""

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Info("Failed to get the file ", err)
		// c.ServeJSON()
		// return
	} else {
		defer file.Close()

		// Save the uploaded file
		fileName := filepath.Base(header.Filename)
		filePath = "/uploads/expense-receipts/" + time.Now().Format("20060102150405") + fileName // Define your file path
		err = c.SaveToFile("ReceiptImage", "../images/"+filePath)
		if err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			// errorMessage := "Error: Failed to save the image file"

			// resp := responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error updating user. " + errorMessage}

			// c.Data["json"] = resp
			// c.ServeJSON()
			// return
		}
	}

	var expenseDate time.Time
	proceed := false
	message := "An error occurred adding this audit request"
	statusCode := 308

	var allowedDateList [4]string = [4]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000"}

	for _, date_ := range allowedDateList {
		logs.Debug("About to convert ", rexpenseDate)
		// Convert dob string to date
		tDate, error := time.Parse(date_, rexpenseDate)

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

		if rcurrency != "" {
			logs.Info("Currency is not empty")
			currencyId, _ := strconv.ParseInt(rcurrency, 10, 64)
			if cur, err := models.GetCurrenciesById(currencyId); err == nil {
				currency = *cur
			}
		}

		category, _ := strconv.ParseInt(rcategory, 10, 64)
		if category, err := models.GetPayment_categoriesById(category); err == nil {
			paymentMethod, _ := strconv.ParseInt(rpaymentmethod, 10, 64)
			if paymentMethod, err := models.GetPayment_methodsById(paymentMethod); err == nil {
				if ruser != "" {
					userId, _ := strconv.ParseInt(ruser, 10, 64)
					if user, err := models.GetUsersById(userId); err == nil {
						amount, _ := strconv.ParseFloat(ramount, 64)
						var branch *string
						if user.UserDetails.Branch != nil {
							branch = &user.UserDetails.Branch.Branch
						}

						expense := models.Expense_records{ExpenseDate: expenseDate, Branch: *branch, Category: category, ReceiptImagePath: filePath, Description: rdescription, Amount: amount, Currency: &currency, PaymentMethod: paymentMethod, Active: 1, CreatedBy: user, ModifiedBy: user, DateCreated: time.Now(), DateModified: time.Now()}

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
						message = "Error adding expense. User provided does not exist"
						statusCode = 608
						resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: nil, StatusDesc: message}
						c.Data["json"] = resp
					}
				} else {
					logs.Info("Error adding expense. User field is empty")
					message = "Error adding expense. User not provided"
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
	message := "An error occurred adding this audit request"
	statusCode := 308
	v, err := models.GetExpense_recordsById(id)
	if err != nil {
		logs.Error("An error occurred fetching expense ", err.Error())
		message = "Error fetching expenses."
		statusCode = 608
		resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = v
		message = "Expense details fetched successfully."
		statusCode = 200
		resp := responses.ExpenseResponse{StatusCode: statusCode, Expense: v, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Expense_records
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Fields returned. e.g. transport ..."
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

	l, err := models.GetAllExpense_records(query, fields, sortby, order, offset, limit, search)
	if err != nil {
		logs.Info("Error fetching expenses ", err.Error())
		message = "Error fetching expenses."
		statusCode = 608
		resp := responses.ExpensesResponseDTO{StatusCode: statusCode, Expenses: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Info("Expenses fetched are ", l)
		if l == nil {
			l = []interface{}{}
		}
		statusCode = 200
		message = "Expenses fetched successfully"
		resp := responses.ExpensesResponseDTO{StatusCode: statusCode, Expenses: &l, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllByBranch ...
// @Title Get All By Branch
// @Description get Expense_records
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Expense_records
// @Failure 403
// @router /branch/:id [get]
func (c *Expense_recordsController) GetAllByBranch() {
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

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	if branch, err := models.GetBranchesById(id); err == nil {
		query = map[string]string{"branch": branch.Branch}
		l, err := models.GetAllExpense_records(query, fields, sortby, order, offset, limit, search)
		if err != nil {
			logs.Info("Error fetching expenses ", err.Error())
			message = "Error fetching expenses."
			statusCode = 608
			resp := responses.ExpensesResponseDTO{StatusCode: statusCode, Expenses: nil, StatusDesc: message}
			c.Data["json"] = resp
		} else {
			if l == nil {
				l = []interface{}{}
			}
			statusCode = 200
			message = "Expenses fetched successfully"
			resp := responses.ExpensesResponseDTO{StatusCode: statusCode, Expenses: &l, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		logs.Info("Error fetching expenses ", err.Error())
		message = "Error fetching expenses."
		statusCode = 608
		resp := responses.ExpensesResponseDTO{StatusCode: statusCode, Expenses: nil, StatusDesc: message}
		c.Data["json"] = resp
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

// GetExpenseRecordCount ...
// @Title Get Expense Record Count
// @Description get Count of expense records
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/ [get]
func (c *Expense_recordsController) GetExpenseRecordCount() {
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

	v, err := models.GetExpenseRecordCount(query, search)
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
