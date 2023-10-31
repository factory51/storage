package engine

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factory51/storage/core/config" //reference to our config package
	"github.com/factory51/storage/core/crud"
	"github.com/factory51/storage/core/responses"
)

func Update(w http.ResponseWriter, r *http.Request) {

	server := config.Get("ident").(string)

	bucket := Bucket{}
	err := json.NewDecoder(r.Body).Decode(&bucket)

	if err != nil {
		msg := fmt.Sprintf("Could not decode resource: %v", err.Error())
		responses.StandandResponseMessage(w, 500, server, msg)
		return
	}

	crud.GenericUpdate(w, bucket, "storage_bucket") //call to generic get method

}
