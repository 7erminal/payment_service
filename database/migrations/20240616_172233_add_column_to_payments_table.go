package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToPaymentsTable_20240616_172233 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToPaymentsTable_20240616_172233{}
	m.Created = "20240616_172233"

	migration.Register("AddColumnToPaymentsTable_20240616_172233", m)
}

// Run the migrations
func (m *AddColumnToPaymentsTable_20240616_172233) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE payments add COLUMN transaction_id int default null after initiated_by")
}

// Reverse the migrations
func (m *AddColumnToPaymentsTable_20240616_172233) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
