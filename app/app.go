package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Start starts the web server
func StartApp() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomerId)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
