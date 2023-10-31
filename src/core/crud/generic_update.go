package crud

import (
	"fmt"
	"net/http" // for json manipulation

	"github.com/factory51/storage/core/config"   //reference to our database package for orm
	"github.com/factory51/storage/core/database" //reference to our database package for orm

	//helper functions
	"github.com/factory51/storage/core/orm"       //reference to our config package       //orm
	"github.com/factory51/storage/core/responses" //reference to our config package//feedback
)

func GenericUpdate[T any](w http.ResponseWriter, structure T, table_name string) {

	server := config.Get("ident").(string)

	err := orm.Update(database.Conn, table_name, structure) //use orm to create

	if err != nil {
		msg := fmt.Sprintf("Cannot Update resource: %v\n", err.Error())
		responses.StandandResponseMessage(w, 500, server, msg)
		return
	}

	feedback := fmt.Sprintf("Successfully Updated [%v]", server)
	responses.StandandResponseMessage(w, 200, server, feedback) //return our response
}
