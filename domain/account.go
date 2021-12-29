package domain

import "github.com/adityapwr/go-banking/errs"

type Account struct {
	Id          string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string
	AccountType string
	Balance     float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
