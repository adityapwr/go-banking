package domain

import (
	"strconv"

	"github.com/adityapwr/banking-lib/errs"
	"github.com/adityapwr/banking-lib/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (? , ? , ? , ? , ? )`
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while inserting rows")
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id")
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	a.Id = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
