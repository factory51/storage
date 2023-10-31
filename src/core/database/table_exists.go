package database

import (
	"database/sql" //go inbuilt SQL package
	"fmt"          //string formatting
)

/*
TableExists allows user to check if a database table exists
[ACCEPTS]

	db *sql.DB as a database connection
	database string which database to check inside
	table string which table to check existance thereof

[RETURNS]

	exists bool denoting if the table is real or not
	err err any error messages generate in performing functionality

*/

func TableExists(db *sql.DB, database string, table string) (exists bool, err error) {

	sql := fmt.Sprintf("SELECT COUNT(1) FROM information_schema.TABLES  WHERE TABLE_SCHEMA = '%v' AND TABLE_TYPE = 'BASE TABLE' AND TABLE_NAME = '%v'", database, table)
	query, err := db.Prepare(sql)

	if err != nil {
		return // well we sure messed up here
	}

	rows, err := query.Query() //get our row

	if err != nil {
		return
	}

	for rows.Next() {
		var cnt int

		err := rows.Scan(&cnt)

		if err != nil {
			return false, err
		}

		if cnt == 1 { //there can only be one
			exists = true
		}
	}

	return
}
