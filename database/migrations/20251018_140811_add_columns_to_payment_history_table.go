package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToPaymentHistoryTable_20251018_140811 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToPaymentHistoryTable_20251018_140811{}
	m.Created = "20251018_140811"

	migration.Register("AddColumnsToPaymentHistoryTable_20251018_140811", m)
}

// Run the migrations
func (m *AddColumnsToPaymentHistoryTable_20251018_140811) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE payment_history ADD COLUMN service varchar(128) DEFAULT NULL AFTER status, ADD COLUMN narration text DEFAULT NULL AFTER service, ADD COLUMN reference varchar(128) default null after narration ;")
}

// Reverse the migrations
func (m *AddColumnsToPaymentHistoryTable_20251018_140811) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
