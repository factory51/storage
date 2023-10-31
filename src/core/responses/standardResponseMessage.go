package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StandandResponseMessage(w http.ResponseWriter, httpCode int, ident string, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Factory51")
	w.Header().Set("X-Server-Ident", ident)
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(httpCode)

	response := new(Response)
	response.Code = httpCode
	response.Message = message

	js, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(js)
	return

}
