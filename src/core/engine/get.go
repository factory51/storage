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

func GetList(w http.ResponseWriter, r *http.Request) {

	server := config.Get("ident").(string)
	fmt.Printf("GetList Ident: %v\n", server)

	key_list := GetKeys()

	payload, err := helpers.StructToJson(key_list)

	if err != nil {
		msg := fmt.Sprintf("Could not decode key list")
		responses.StandandResponseMessage(w, 500, server, msg)
		return
	}

	responses.StandandResponseJson(w, 200, server, payload) //return our response
	return

}

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

func GetKeys() (key_list []KeyItem) {

	sql := fmt.Sprintf("SELECT `key` FROM storage_bucket")

	query, err := database.Conn.Prepare(sql)

	if err != nil {
		fmt.Printf("[FATAL] Cannot structure GetKeys() query correctly: %v\n", err.Error())
	}

	rows, err := query.Query()

	for rows.Next() {

		keyItem := KeyItem{}

		err := rows.Scan(&keyItem.Key)

		if err != nil {
			fmt.Printf("[FATAL] Cannot Scan GetKeys() query correctly: %v\n", err.Error())
		} else {
			key_list = append(key_list, keyItem)
		}

	}

	return

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
