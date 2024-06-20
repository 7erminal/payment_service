package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PaymentTypes_20240523_183803 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PaymentTypes_20240523_183803{}
	m.Created = "20240523_183803"

	migration.Register("PaymentTypes_20240523_183803", m)
}

// Run the migrations
func (m *PaymentTypes_20240523_183803) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE payment_types(`payment_type_id` int(11) NOT NULL AUTO_INCREMENT,`payment_type` varchar(255) NOT NULL,`description` varchar(255) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`payment_type_id`))")
}

// Reverse the migrations
func (m *PaymentTypes_20240523_183803) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `payment_types`")
}
