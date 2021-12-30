package service

import (
	"time"

	"github.com/adityapwr/go-banking/domain"
	"github.com/adityapwr/go-banking/dto"
	"github.com/adityapwr/go-banking/errs"
)

type AccountService interface {
	AddNewAccount(dto.NewAccountRequest) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) AddNewAccount(a dto.NewAccountRequest) (*dto.AccountResponse, *errs.AppError) {
	account, err := s.repo.Save(domain.Account{
		CustomerId:  a.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      "1",
	})
	if err != nil {
		return nil, err
	}
	response := account.ToDto()
	return &response, nil

}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
