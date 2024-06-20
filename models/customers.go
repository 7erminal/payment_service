package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Customers struct {
	CustomerId       int64                `orm:"auto"`
	User             *Users               `orm:"rel(fk)"`
	CustomerCategory *Customer_categories `orm:"rel(fk);omitempty;null"`
	Nickname         string               `orm:"size(100);omitempty;null"`
	DateCreated      time.Time            `orm:"type(datetime)"`
	DateModified     time.Time            `orm:"type(datetime)"`
	CreatedBy        int
	ModifiedBy       int
	Active           int
}

func init() {
	orm.RegisterModel(new(Customers))
}
