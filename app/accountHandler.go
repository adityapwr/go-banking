package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adityapwr/go-banking/dto"
	"github.com/adityapwr/go-banking/logger"
	"github.com/adityapwr/go-banking/service"
	"github.com/gorilla/mux"
)

type Account struct {
	CustomerId  string `json:"customer_id"`
	OpeningDate string
	AccountType string
	Balance     float64
	Status      string
}

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) AddNewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var account dto.NewAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		writeResponse(w, http.StatusInternalServerError, nil)
	}
	account.CustomerId = customerId
	logger.Info(fmt.Sprintf("Adding new account for customer %s", account.CustomerId))
	newAccount, err := ah.service.AddNewAccount(account)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, nil)
	}
	writeResponse(w, http.StatusOK, newAccount)
}
