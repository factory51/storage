package engine

import (
	"fmt"
	"net/http"

	"github.com/factory51/storage/core/config"   //reference to our config package
	"github.com/factory51/storage/core/database" //reference to our database package
	"github.com/factory51/storage/core/helpers"
	"github.com/factory51/storage/core/responses" //reference to our responses package

	"github.com/gorilla/mux" //gorilla
)

func Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ident := vars["ident"]
	server := config.Get("ident").(string)

	fmt.Printf("Get Ident: %v\n", ident)

	bucket, found := GetPiece(ident)

	if found {

		payload, err := helpers.StructToJson(bucket)

		if err != nil {
			msg := fmt.Sprintf("Could not decode resource: %v", ident)
			responses.StandandResponseMessage(w, 500, server, msg)
			return
		}

		responses.StandandResponseJson(w, 200, server, payload) //return our response
		return

	} else {

		msg := fmt.Sprintf("Could not find resource: %v", ident)
		responses.StandandResponseMessage(w, 404, server, msg)
		return

	}
}

func GetPiece(ident string) (bucket Bucket, found bool) {

	sql := fmt.Sprintf("SELECT * FROM storage_bucket WHERE `key`='%v'", ident)

	query, err := database.Conn.Prepare(sql)

	if err != nil {
		fmt.Printf("[FATAL] Cannot structure GetPiece(%v) query correctly: %v\n", ident, err.Error())
	}

	rows, err := query.Query()

	for rows.Next() {

		err := rows.Scan(&bucket.CreateAt, &bucket.ClientIdent, &bucket.Key, &bucket.Value)

		if err != nil {
			fmt.Printf("[FATAL] Cannot Scan GetPiece(%v) query correctly: %v\n", ident, err.Error())
		} else {
			found = true
		}

	}
	return
}
