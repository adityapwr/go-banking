package app

import (
	"log"
	"net/http"

	"github.com/adityapwr/go-banking/domain"
	"github.com/adityapwr/go-banking/service"
	"github.com/gorilla/mux"
)

//Start starts the web server
func StartApp() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
