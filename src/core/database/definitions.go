package database

import (
	"database/sql" //sql lib
)

var Conn *sql.DB                                  //our database connection
const connection_skeleton = "%s:%s@tcp(%s:%s)/%s" //skeleton structure for our database connection string

type ConnectionDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
} // structure to hold connection details
