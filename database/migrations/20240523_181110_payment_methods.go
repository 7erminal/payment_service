package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PaymentMethods_20240523_181110 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PaymentMethods_20240523_181110{}
	m.Created = "20240523_181110"

	migration.Register("PaymentMethods_20240523_181110", m)
}

// Run the migrations
func (m *PaymentMethods_20240523_181110) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE payment_methods(`payment_method_id` int(11) NOT NULL AUTO_INCREMENT,`payment_method` varchar(128) NOT NULL,`description` varchar(255) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`payment_method_id`))")
}

// Reverse the migrations
func (m *PaymentMethods_20240523_181110) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `payment_methods`")
}
