package controllers

import (
	"encoding/json"
	"payment_service/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

// AccountsController operations for Accounts
type AccountsController struct {
	beego.Controller
}

// URLMapping ...
func (c *AccountsController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
}

// GetOne ...
// @Title Get One
// @Description get Accounts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAccountsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Accounts	true		"body for Accounts content"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AccountsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Accounts{AccountId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateAccountsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
