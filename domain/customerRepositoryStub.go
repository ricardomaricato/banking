package domain

import "github.com/ricardomaricato/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(string) (*Customer, *errs.AppError) {
	return nil, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
