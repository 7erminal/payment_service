package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToPaymentsTable_20251013_191807 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToPaymentsTable_20251013_191807{}
	m.Created = "20251013_191807"

	migration.Register("AddColumnToPaymentsTable_20251013_191807", m)
}

// Run the migrations
func (m *AddColumnToPaymentsTable_20251013_191807) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE payments ADD COLUMN request_id int AFTER transaction_id, ADD COLUMN commission float DEFAULT 0 AFTER amount, ADD COLUMN charge float DEFAULT 0 AFTER commission;")
	m.SQL("ALTER TABLE payments ADD CONSTRAINT fk_payments_request_id FOREIGN KEY (request_id) REFERENCES requests(request_id) ON UPDATE CASCADE ON DELETE NO ACTION;")
	m.SQL("ALTER TABLE payments ADD COLUMN narration varchar(255) DEFAULT NULL AFTER charge;")
}

// Reverse the migrations
func (m *AddColumnToPaymentsTable_20251013_191807) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
