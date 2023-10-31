package database

import (
	"database/sql" //go inbuilt SQL package
	"log"
	"os"
	"strings"
)

/*
	TransactionFromFile runs a set of sql statments as a transaction specified in the file thats referenced, multiple statements can be run in one file but must be terminated with ;

[Accepts]

	db *sql.DB as a database connection
	path string path to the file with the sql statment(s)
*/
func TransactionFromFile(db *sql.DB, path string) (err error) {

	err = db.Ping()

	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("db.Begin() ERR: %v\n", err)
		return err
	}

	defer func() {
		tx.Rollback()
	}()

	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}

		if _, err := tx.Exec(q); err != nil {
			log.Printf("tx.Exec(q) ERR: %v\n", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Cannot perfrom transaction: %v\n", err.Error())
	}

	return
}
