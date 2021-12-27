package domain

import "github.com/adityapwr/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Pincode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
