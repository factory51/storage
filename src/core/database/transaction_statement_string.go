package database

import (
	"database/sql" //go inbuilt SQL package

	"strings"
)

/*
	TransactionFromString runs a set of sql statments as a transaction specified in the file thats referenced, multiple statements can be run in one string but must be terminated with ;

[Accepts]

	db *sql.DB as a database connection
	path string path to the file with the sql statment(s)
*/

func TransactionFromString(db *sql.DB, statement string) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Rollback()
	}()
	for _, q := range strings.Split(string(statement), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	return tx.Commit()
}
