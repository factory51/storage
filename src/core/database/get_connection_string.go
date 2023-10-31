package database

import (
	"fmt" //used to format string
)

/*
	GetConnectionString creates a formatted connection based on the supplied input

[ACCEPTS]

	user string username to use to connect to database with
	pass string password to use to connect to database with
	host string host to use to connect to database with
	port int port to use to connect to database with
	database string database to connect to

[RETURNS]

	connectionString string the formatted connection string
*/
func GetConnectionString(user string, pass string, host string, port string, database string) (connectionString string) {

	connectionString = fmt.Sprintf(connection_skeleton, user, pass, host, port, database)
	return
}
