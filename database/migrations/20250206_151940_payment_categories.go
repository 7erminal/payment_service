package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PaymentCategories_20250206_151940 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PaymentCategories_20250206_151940{}
	m.Created = "20250206_151940"

	migration.Register("PaymentCategories_20250206_151940", m)
}

// Run the migrations
func (m *PaymentCategories_20250206_151940) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE payment_categories(`payment_category_id` int(11) NOT NULL AUTO_INCREMENT,`category` varchar(80) NOT NULL,`description` varchar(255) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`payment_category_id`))")
}

// Reverse the migrations
func (m *PaymentCategories_20250206_151940) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `payment_categories`")
}
