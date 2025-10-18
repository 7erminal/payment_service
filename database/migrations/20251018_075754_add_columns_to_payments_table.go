package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToPaymentsTable_20251018_075754 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToPaymentsTable_20251018_075754{}
	m.Created = "20251018_075754"

	migration.Register("AddColumnsToPaymentsTable_20251018_075754", m)
}

// Run the migrations
func (m *AddColumnsToPaymentsTable_20251018_075754) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`ALTER TABLE payments ADD COLUMN other_charge DOUBLE DEFAULT 0 AFTER charge, ADD COLUMN date_processed datetime DEFAULT null AFTER date_modified, ADD COLUMN payment_amount DOUBLE DEFAULT 0 AFTER payment_method, ADD COLUMN service varchar(128) DEFAULT null AFTER sender;`)

}

// Reverse the migrations
func (m *AddColumnsToPaymentsTable_20251018_075754) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
