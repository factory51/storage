package orm

import (

	//go inbuilt SQL package
	"database/sql"

	//visual debugging
	_ "gorm.io/driver/mysql" //gorm mysql package
)

/*
Create generic function to perform a database INSERT INTO supplied table and data structure

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	record T generics reference to the data structure thats going to be created

[Returns]

	record T amended record after the create has happened. Contains structure specific id (so this function is free from differing structure id names)
	err error any generated errors
*/
func Create[T any](conn *sql.DB, table string, record *T) (err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Create(&record)

	if results.Error != nil {
		err = results.Error
	}
	return

}
