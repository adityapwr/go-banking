package domain

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}

func NewCustomerrepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Aditya", City: "Bangalore", Pincode: "560037", DateofBirth: "01/01/1990", Status: "Active"}}
	return CustomerRepositoryStub{Customers: customers}
}
