package routing

import (
	"fmt"      //string formatting
	"net/http" // used for http control

	"github.com/factory51/storage/core/engine"     //reference to our engine / handlers package
	"github.com/factory51/storage/core/middleware" //reference to our middleware package

	"github.com/gorilla/handlers" // gorilla handler functions to get around pesky CORS dumbfuckery
	"github.com/gorilla/mux"      //gorilla mux router for our routing.
)

func HandleRoutes(port int) {

	//our gorilla mux router
	router := mux.NewRouter()

	storageRoutes(router)

	//setting up CORS exceltpion headers
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//setting up CORS exceltpion Origins: set to all origins at the moment as assumption is middleware will handle IP restrictions
	originsOk := handlers.AllowedOrigins([]string{"*"})
	//setting up CORS exceltpion Methods: set to all Methods at the moment as assumption is middleware will handle HTTP method restrictions
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//follwings is some useful output so you can see what url the api is being served on
	full_domain := fmt.Sprintf(":%v", port) //lock to a port not a url and let ngnix direct traffic

	//listen and serve biatches with our CORS exceptions added to the router
	http.ListenAndServe(full_domain, handlers.CORS(originsOk, headersOk, methodsOk)(router))

}

func storageRoutes(r *mux.Router) {

	vmw := middleware.ChainMiddleware(middleware.Verify) //protected by middleware

	//r.Handle("/{mode}/account/{action}", controllers.AccountDispatcher)  //application login
	//r.Handle("/{mode}/account/{action}/", controllers.AccountDispatcher) //trailing slash catchall

	r.HandleFunc("/get/{ident}", vmw(engine.Get)).Methods("GET")
	r.HandleFunc("/get", vmw(engine.GetList)).Methods("GET")
	r.HandleFunc("/create", vmw(engine.Create)).Methods("POST")
	r.HandleFunc("/update", vmw(engine.Update)).Methods("PUT")
	r.HandleFunc("/docs", engine.Documentation).Methods("GET")

}
