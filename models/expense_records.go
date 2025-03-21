package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Expense_records struct {
	ExpenseRecordId  int64               `orm:"auto"`
	Category         *Payment_categories `orm:"rel(fk); column(category)"`
	ReceiptImagePath string
	Branch           string
	Description      string           `orm:"size(255)"`
	Currency         *Currencies      `orm:"rel(fk);column(currency_id)"`
	PaymentMethod    *Payment_methods `orm:"rel(fk);column(payment_method_id)"`
	Amount           float64
	ExpenseDate      time.Time `orm:"type(datetime)"`
	Active           int
	DateCreated      time.Time `orm:"type(datetime)"`
	DateModified     time.Time `orm:"type(datetime)"`
	CreatedBy        *Users    `orm:"rel(fk);column(created_by)"`
	ModifiedBy       *Users    `orm:"rel(fk);column(modified_by)"`
}

func init() {
	orm.RegisterModel(new(Expense_records))
}

// AddExpense_records insert a new Expense_records into database and returns
// last inserted Id on success.
func AddExpense_records(m *Expense_records) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExpense_recordsById retrieves Expense_records by Id. Returns error if
// Id doesn't exist
func GetExpense_recordsById(id int64) (v *Expense_records, err error) {
	o := orm.NewOrm()
	v = &Expense_records{ExpenseRecordId: id}
	if err = o.QueryTable(new(Expense_records)).Filter("ExpenseRecordId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetOrderCount retrieves Items by Id. Returns error if
// Id doesn't exist
func GetExpenseRecordCount(query map[string]string, search map[string]string) (c int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Expense_records))
	if len(search) > 0 {
		cond := orm.NewCondition()
		for k, v := range search {
			// rewrite dot-notation to Object__Attribute
			k = strings.Replace(k, ".", "__", -1)
			if strings.Contains(k, "isnull") {
				qs = qs.Filter(k, (v == "true" || v == "1"))
			} else {
				logs.Info("Adding or statement")
				cond = cond.Or(k+"__icontains", v)

				// qs = qs.Filter(k+"__icontains", v)

			}
		}
		logs.Info("Condition set ", qs)
		qs = qs.SetCond(cond)
	}
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
		logs.Info("Modified query is ", k, " and ", v)
	}
	if c, err = qs.RelatedSel().Count(); err == nil {
		return c, nil
	}
	return 0, err
}

// GetOrderCount retrieves Items by Id. Returns error if
// Id doesn't exist
func GetExpenseRecordCountByType() (c int64, err error) {
	o := orm.NewOrm()
	if c, err = o.QueryTable(new(Expense_records)).Count(); err == nil {
		return c, nil
	}
	return 0, err
}

// GetAllExpense_records retrieves all Expense_records matches certain condition. Returns empty list if
// no records exist
func GetAllExpense_records(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, search map[string]string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Expense_records))
	if len(search) > 0 {
		cond := orm.NewCondition()
		for k, v := range search {
			// rewrite dot-notation to Object__Attribute
			k = strings.Replace(k, ".", "__", -1)
			if strings.Contains(k, "isnull") {
				qs = qs.Filter(k, (v == "true" || v == "1"))
			} else {
				logs.Info("Adding or statement")
				cond = cond.Or(k+"__icontains", v)

				// qs = qs.Filter(k+"__icontains", v)

			}
		}
		logs.Info("Condition set ", qs)
		qs = qs.SetCond(cond)
	}

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Expense_records
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// GetAllExpense_records retrieves all Expense_records matches certain condition. Returns empty list if
// no records exist
func GetAllExpense_recordsByBranch(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, branch string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Expense_records))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Expense_records
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateExpense_records updates Expense_records by Id and returns error if
// the record to be updated doesn't exist
func UpdateExpense_recordsById(m *Expense_records) (err error) {
	o := orm.NewOrm()
	v := Expense_records{ExpenseRecordId: m.ExpenseRecordId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExpense_records deletes Expense_records by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExpense_records(id int64) (err error) {
	o := orm.NewOrm()
	v := Expense_records{ExpenseRecordId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Expense_records{ExpenseRecordId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
