package domain

import (
	"github.com/adityapwr/go-banking/dto"
	"github.com/adityapwr/go-banking/errs"
)

type Account struct {
	Id          string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToDto() dto.AccountResponse {
	response := dto.AccountResponse{
		Id: a.Id,
	}
	return response
}
