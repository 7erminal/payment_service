package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Customer_categories struct {
	CustomerCategoryId int64     `orm:"auto"`
	Category           string    `orm:"size(100)"`
	Description        string    `orm:"size(255); null"`
	DateCreated        time.Time `orm:"type(datetime)"`
	DateModified       time.Time `orm:"type(datetime)"`
	CreatedBy          int
	ModifiedBy         int
	Active             int
}

func init() {
	orm.RegisterModel(new(Customer_categories))
}
