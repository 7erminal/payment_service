package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	UserId        int64             `orm:"column(user_id);auto"`
	UserDetails   *UserExtraDetails `orm:"rel(fk);column(user_details_id);null"`
	ImagePath     string            `orm:"column(image_path);size(255);null"`
	UserType      int               `orm:"column(user_type);null"`
	FullName      string            `orm:"column(full_name);size(255)"`
	Username      string            `orm:"column(username);size(40);null"`
	Password      string            `orm:"column(password);size(255)"`
	Email         string            `orm:"column(email);size(255);null"`
	PhoneNumber   string            `orm:"column(phone_number);size(255);null"`
	Gender        string            `orm:"column(gender);size(10)"`
	Dob           time.Time         `orm:"column(dob);type(datetime)"`
	Address       string            `orm:"column(address);size(255);null"`
	IdType        string            `orm:"column(id_type);size(5);null"`
	IdNumber      string            `orm:"column(id_number);size(100);null"`
	MaritalStatus string            `orm:"column(marital_status);size(20);null"`
	Active        int               `orm:"column(active);null"`
	IsVerified    bool              `orm:"column(is_verified);null"`
	DateCreated   time.Time         `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified  time.Time         `orm:"column(date_modified);type(datetime);null"`
	CreatedBy     int               `orm:"column(created_by);null"`
	ModifiedBy    int               `orm:"column(modified_by);null"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// GetUsersById retrieves Users by username. Returns error if
// Id doesn't exist
func GetUsersByUsername(username string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Email: username}
	if err = o.QueryTable(new(Users)).Filter("Email", username).RelatedSel().One(v); err == nil {
		return v, nil
	} else if err = o.QueryTable(new(Users)).Filter("PhoneNumber", username).RelatedSel().One(v); err == nil {
		return v, nil
	} else if err = o.QueryTable(new(Users)).Filter("Username", username).RelatedSel().One(v); err == nil {
		return v, nil
	}

	return nil, err
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int64) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{UserId: id}
	if err = o.QueryTable(new(Users)).Filter("UserId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
