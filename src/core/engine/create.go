package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/factory51/storage/core/config" //reference to our config package
	"github.com/factory51/storage/core/crud"
	"github.com/factory51/storage/core/responses" //reference to our responses package
)

func Create(w http.ResponseWriter, r *http.Request) {

	server := config.Get("ident").(string)

	bucket := Bucket{}
	err := json.NewDecoder(r.Body).Decode(&bucket)

	if err != nil {
		msg := fmt.Sprintf("Could not decode resource: %v", err.Error())
		responses.StandandResponseMessage(w, 500, server, msg)
		return
	}

	currentTime := time.Now()

	timestamp := currentTime.Format("2006-01-02 15:04:05")

	bucket.CreateAt = timestamp

	crud.GenericCreate(w, r, bucket, "storage_bucket") //call to generic create

}
