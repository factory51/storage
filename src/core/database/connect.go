package database

import (
	"database/sql" //go inbuilt SQL package
	"log"

	_ "github.com/go-sql-driver/mysql" //go mysql mySQL, indirect reference
)

// Connect using supplied connection string to establish a database connection
func Connect(connection_string string) (err error) {

	Conn, err = sql.Open("mysql", connection_string) //connect to mysql

	if err != nil {
		log.Printf("[CORE ERROR] cannot connect to %v\nReason: %s\n", connection_string, err.Error())
	}

	return
}
