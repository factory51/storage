package responses

import (
	"net/http"
)

func StandandResponseJson(w http.ResponseWriter, httpCode int, ident string, payload string) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Factory51")
	w.Header().Set("X-Server-Ident", ident)
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(httpCode)

	w.Write([]byte(payload))

}
