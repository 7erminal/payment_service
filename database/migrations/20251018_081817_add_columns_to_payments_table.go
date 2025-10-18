package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToPaymentsTable_20251018_081817 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToPaymentsTable_20251018_081817{}
	m.Created = "20251018_081817"

	migration.Register("AddColumnsToPaymentsTable_20251018_081817", m)
}

// Run the migrations
func (m *AddColumnsToPaymentsTable_20251018_081817) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`ALTER TABLE payments ADD COLUMN sender_account varchar(128) DEFAULT NULL AFTER reciever, ADD COLUMN receiver_account varchar(128) DEFAULT NULL AFTER sender_account;`)

}

// Reverse the migrations
func (m *AddColumnsToPaymentsTable_20251018_081817) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
