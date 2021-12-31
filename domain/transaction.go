package domain

import (
	"github.com/adityapwr/go-banking/dto"
	"github.com/adityapwr/go-banking/errs"
)

const WITHDRAWAL = "WITHDRAWAL"

type Transaction struct {
	Id              string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	Date            string  `db:"transaction_date"`
}

type TransactionRepository interface {
	Transaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
}

func (t Transaction) ToDto() dto.TransactionResponse {
	response := dto.TransactionResponse{
		TransactionId: t.Id,
		Amount:        t.Amount,
	}
	return response
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
