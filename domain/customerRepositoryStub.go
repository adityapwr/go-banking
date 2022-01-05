package domain

import "github.com/adityapwr/banking-lib/errs"

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll(_ string) ([]Customer, *errs.AppError) {
	return s.Customers, nil
}

func (s CustomerRepositoryStub) ById(_ string) (*Customer, *errs.AppError) {
	customer := Customer{
		Id:          "1",
		Name:        "Aditya",
		City:        "Bangalore",
		Pincode:     "560037",
		DateofBirth: "01/01/1990",
		Status:      "Active",
	}
	return &customer, nil
}

func NewCustomerrepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Aditya", City: "Bangalore", Pincode: "560037", DateofBirth: "01/01/1990", Status: "Active"}}
	return CustomerRepositoryStub{Customers: customers}
}
