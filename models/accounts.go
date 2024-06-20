package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Accounts struct {
	AccountId     int64 `orm:"auto"`
	UserId        int
	AccountNumber string `orm:"size(255)"`
	Balance       float64
	BalanceBefore float64
	DateCreated   time.Time `orm:"type(datetime)"`
	DateModified  time.Time `orm:"type(datetime)"`
	CreatedBy     int
	ModifiedBy    int
	Active        int
}

func init() {
	orm.RegisterModel(new(Accounts))
}

// GetAccountsById retrieves Accounts by Id. Returns error if
// Id doesn't exist
func GetAccountsById(id int64) (v *Accounts, err error) {
	o := orm.NewOrm()
	v = &Accounts{AccountId: id}
	if err = o.QueryTable(new(Accounts)).Filter("AccountId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateAccounts updates Accounts by Id and returns error if
// the record to be updated doesn't exist
func UpdateAccountsById(m *Accounts) (err error) {
	o := orm.NewOrm()
	v := Accounts{AccountId: m.AccountId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
