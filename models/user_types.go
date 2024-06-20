package models

import (
	"time"
)

type UserTypes struct {
	Id                  int       `orm:"column(user_type_id);auto"`
	UserTypeName        string    `orm:"column(user_type_name);size(255)"`
	UserTypeDescription string    `orm:"column(user_type_description);size(255)"`
	Active              int       `orm:"column(active);null"`
	DateCreated         time.Time `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified        time.Time `orm:"column(date_modified);type(datetime);null"`
	CreatedBy           int       `orm:"column(created_by);null"`
	ModifiedBy          int       `orm:"column(modified_by);null"`
}

func (t *UserTypes) TableName() string {
	return "user_types"
}
