package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToPaymentsTable_20251019_201303 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToPaymentsTable_20251019_201303{}
	m.Created = "20251019_201303"

	migration.Register("AddColumnsToPaymentsTable_20251019_201303", m)
}

// Run the migrations
func (m *AddColumnsToPaymentsTable_20251019_201303) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE payments ADD COLUMN payment_currency varchar(5) DEFAULT NULL AFTER payment_amount;")

}

// Reverse the migrations
func (m *AddColumnsToPaymentsTable_20251019_201303) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
