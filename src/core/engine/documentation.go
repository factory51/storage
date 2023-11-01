package engine

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	file := "./html/index.html"

	http.ServeFile(w, r, file)

}

func Documentation(w http.ResponseWriter, r *http.Request) {

	file := "./html/docs.html"

	http.ServeFile(w, r, file)

}
