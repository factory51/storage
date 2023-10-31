package orm

import (
	"database/sql" //go inbuilt SQL package
)

/*
Delete generic function to perform a database UPDATE based on supplied table and data structure

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	record T generics reference to the data structure thats going to be updated

[Returns]

	record T amended record after the create has happened. Contains structure specific id (so this function is free from differing structure id names)
	err error any generated errors
*/
func Delete[T any](conn *sql.DB, table string, example T) (err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Delete(&example)

	if results.Error != nil {
		err = results.Error
	}
	return

}
