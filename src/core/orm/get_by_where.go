package orm

import (

	//go inbuilt SQL package
	"database/sql"

	//visual debugging
	_ "gorm.io/driver/mysql" //gorm mysql package
)

/*
GetByWhere generic function to perform a database SELECT * FROM ... with where and value being used to build the where clause.

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	example T generics reference to the data structure thats going to be created
	where string holds the basis of the where clause gorm will build up eg field_name = ?
	value string the value for gorm to replace into the ? holder in there where string

[Returns]

	object []T slice of supplied example T
	err error any generated errors
*/
func GetByWhere[T any](conn *sql.DB, table string, example T, where string, value string) (objects []T, err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Where(where, value).Find(&objects)

	if results.Error != nil {
		err = results.Error
	}
	return

}
