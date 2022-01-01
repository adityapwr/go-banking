package domain

import (
	"github.com/adityapwr/banking-lib/errs"
	"github.com/adityapwr/go-banking/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Pincode     string `db:"zipcode"`
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusText() string {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}

func (c Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Pincode:     c.Pincode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusText(),
	}
	return response
}
