package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adityapwr/go-banking/domain"
	"github.com/adityapwr/go-banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("SERVER_ADDRESS not set")
	}
}

//Start starts the web server
func StartApp() {
	sanityCheck()
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	transactionRepositoryDb := domain.NewTransactionRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandlers{service: service.NewAccountService(accountRepositoryDb)}
	th := TransactionHandlers{service: service.NewTransactionService(transactionRepositoryDb)}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/AddNewAccount", ah.AddNewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{account_id:[0-9]+}/transaction", th.Withdraw).Methods(http.MethodPost)
	SERVER_ADDRESS := os.Getenv("SERVER_ADDRESS")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", SERVER_ADDRESS, SERVER_PORT), router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section...
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
