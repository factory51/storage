package main

import (
	"flag"
	"fmt"

	"github.com/factory51/storage/core/config"   //reference to our config package
	"github.com/factory51/storage/core/database" //reference to our database package
	"github.com/factory51/storage/core/routing"  //reference to our routes package
)

var config_path string //console flag from the path to the config file to load to bootstrap the application
var port int           //Console flag for overriding listening port

func init() {

	fmt.Printf("Initializing Storage Soultions\n")

	flag.StringVar(&config_path, "config", "./conf/app.conf.json", "-config /path/to/config.file") //config cli param
	flag.IntVar(&port, "port", 8081, "-port {NUMBER")                                              //default port to run on is 8081
	flag.Parse()

	/*
		Load our supplied config file. If not supplied defaults to ./conf/app.cong.json
	*/

	file_path, file_name, file_type := config.ParseViperConfigPath(config_path)
	err := config.LoadConfig(file_path, file_name, file_type)

	if err != nil { //we need the config file to setup logs, one and only time error is handled outside the handlers
		fmt.Printf("[Error]: Cannot Open and Load application config file:\n%v\n", err.Error())
	}

	username := config.Get("database.username").(string)
	password := config.Get("database.password").(string)
	host := config.Get("database.host").(string)
	port := config.Get("database.port").(string)
	db := config.Get("database.database").(string)

	dbConnStr := database.GetConnectionString(username, password, host, port, db) //format our connection
	err = database.Connect(dbConnStr)

	if err != nil { //we need the config file to setup logs, one and only time error is handled outside the handlers
		fmt.Printf("[Error]: Cannot Connect to Database\n%v\n", err.Error())
	}
}

func main() {

	fmt.Printf("Starting Storage Soultions\n")
	routing.HandleRoutes(port)

}
