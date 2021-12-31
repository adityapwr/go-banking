package service

import (
	"time"

	"github.com/adityapwr/go-banking/domain"
	"github.com/adityapwr/go-banking/dto"
	"github.com/adityapwr/go-banking/errs"
	"github.com/adityapwr/go-banking/logger"
)

type TranscationService interface {
	WithdrawTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	// DepositTransaction(dto.DepositRequest) (*dto.DepositResponse, *errs.AppError)
}

type DefaultTranscationService struct {
	repo domain.TransactionRepository
	// account domain.AccountRepository
}

func (s DefaultTranscationService) WithdrawTransaction(a dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	logger.Info("Withdraw Transaction Start...")
	accountDetails, err := s.repo.FindBy(a.AccountId)
	if err != nil {
		return nil, errs.NewUnexpectedError("Insufficient balance")
	}
	logger.Info("Checking Transaction is possible...")
	if !isTransactionPossible(accountDetails, a.Amount) {
		return nil, errs.NewUnexpectedError("Insufficient balance")
	}
	// logger.Info("Withdrawing amount from account" + strconv.FormatInt(int64(accountDetails.Amount), 10))

	transaction, err := s.repo.Transaction(domain.Transaction{
		AccountId:       a.AccountId,
		Amount:          a.Amount,
		TransactionType: a.TransactionType,
		Date:            time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return nil, err
	}
	response := transaction.ToDto()
	return &response, nil
}

func isTransactionPossible(account *domain.Account, amount float64) bool {
	if account.Amount < amount {
		return false
	}
	return true
}

// func (s DefaultAccountService) DepositTransaction(a dto.DepositRequest) (*dto.DepositResponse, *errs.AppError) {
// 	account, err := s.repo.Deposit(a.AccountId, a.Amount)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := account.ToDto()
// 	return &response, nil
// }

func NewTransactionService(repository domain.TransactionRepository) DefaultTranscationService {
	return DefaultTranscationService{repo: repository}
}
