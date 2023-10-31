package middleware

import (
	"net/http"

	"github.com/factory51/storage/core/config"    //reference to our config package
	"github.com/factory51/storage/core/responses" //reference to our responses package
	//gorilla
)

func Verify(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		api_key := r.Header.Get("authorization") //get the authorization header

		secret := config.Get("app_secret_key").(string)
		server := config.Get("ident").(string)

		if len(api_key) == 0 {
			responses.StandandResponseMessage(w, 403, server, "Not Authorized.")
			return
		}

		if api_key == secret { //hooray our super secure checks have been passed
			next.ServeHTTP(w, r)
		} else { //oh no they haven't
			responses.StandandResponseMessage(w, 403, server, "Not Authorized.")
			return
		}

	}

}
