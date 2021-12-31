package app

import (
	"encoding/json"
	"net/http"

	"github.com/adityapwr/go-banking/dto"
	"github.com/adityapwr/go-banking/service"
	"github.com/gorilla/mux"
)

type TransactionHandlers struct {
	service service.TranscationService
}

func (th TransactionHandlers) Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	account_id := vars["account_id"]
	var transaction dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		// writeResponse(w, err.Code, err.Message)
		writeResponse(w, http.StatusInternalServerError, nil)
	}
	transaction.AccountId = account_id
	reponse, err := th.service.WithdrawTransaction(transaction)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	}
	writeResponse(w, http.StatusOK, reponse)
}
