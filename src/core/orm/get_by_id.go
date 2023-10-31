package orm

import (

	//go inbuilt SQL package
	"database/sql"

	//visual debugging
	_ "gorm.io/driver/mysql" //gorm mysql package
)

/*
GetByID generic function to perform a database SELECT * FROM ... based on the supplied ident

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	example T generics reference to the data structure thats going to be created
	ident int reference to the table identifier (should have gorm reference in datastructure)

[Returns]

	object []T slice of supplied example T
	err error any generated errors
*/
func GetByID[T any](conn *sql.DB, table string, example T, ident int) (objects []T, err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Find(&objects, ident)

	if results.Error != nil {
		err = results.Error
	}
	return

}
