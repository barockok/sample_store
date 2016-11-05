
package main

import (
	"log"
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20161030214852(txn *sql.Tx) {
	sqlStatement := "CREATE TABLE `products` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`name` varchar(255) DEFAULT NULL," +
		"`price` int(11) DEFAULT NULL," +
		"`description` TEXT DEFAULT NULL," +
	  "PRIMARY KEY (`id`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	_, err := txn.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20161030214852(txn *sql.Tx) {
	sqlStatement := "DROP TABLE `products`"
	_, err := txn.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

}
