package domain

import (
	"strconv"

	"github.com/adityapwr/banking-lib/errs"
	"github.com/adityapwr/banking-lib/logger"
	"github.com/jmoiron/sqlx"
)

type transcationRepositoryDb struct {
	client *sqlx.DB
}

func (d transcationRepositoryDb) Transaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	logTranscationQuery := `INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`
	result, _ := tx.Exec(logTranscationQuery, t.AccountId, t.Amount, t.TransactionType, t.Date)
	if t.IsWithdrawal() {
		sqlWithdraw := `UPDATE accounts SET amount = amount - ? WHERE account_id = ?`
		_, err = tx.Exec(sqlWithdraw, t.Amount, t.AccountId)

	} else {
		sqlWithdraw := `UPDATE accounts SET amount = amount + ? WHERE account_id = ?`
		_, err = tx.Exec(sqlWithdraw, t.Amount, t.AccountId)
	}
	if err != nil {
		tx.Rollback()
		logger.Error("Error while withdrawing")
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction")
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	txnId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.Id = strconv.FormatInt(txnId, 10)
	t.Amount = account.Amount

	return &t, nil

}

func (d transcationRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	var accountDetails Account
	logger.Info("Fetching account details for account id: " + accountId)
	sqlQuery := `SELECT customer_id, account_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?`
	err := d.client.Get(&accountDetails, sqlQuery, accountId)
	if err != nil {
		logger.Error("Error while Retriving account")
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	logger.Info("Account details fetched successfully")
	return &accountDetails, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepository {
	return transcationRepositoryDb{dbClient}
}
