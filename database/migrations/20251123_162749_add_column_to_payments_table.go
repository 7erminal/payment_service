package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToPaymentsTable_20251123_162749 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToPaymentsTable_20251123_162749{}
	m.Created = "20251123_162749"

	migration.Register("AddColumnToPaymentsTable_20251123_162749", m)
}

// Run the migrations
func (m *AddColumnToPaymentsTable_20251123_162749) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE payments ADD COLUMN service_network VARCHAR(128) DEFAULT NULL AFTER payment_account, ADD COLUMN service_package VARCHAR(128) DEFAULT NULL AFTER service_network;")
}

// Reverse the migrations
func (m *AddColumnToPaymentsTable_20251123_162749) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
