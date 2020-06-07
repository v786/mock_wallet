package customer

import "fmt"

//Customer : End customer that will use the wallet
type Customer struct {
	Name        string
	Balance     int
	CreditLimit int
}

func (customer *Customer) Create(name string, creditLimit int) error {
	customer.Name = name
	customer.CreditLimit = creditLimit
	customer.Balance = creditLimit

	fmt.Println("customer Created:", customer.Name, customer.CreditLimit)

	return nil
}

func (customer *Customer) GetDetails() string {
	details := fmt.Sprintf("Name: %s\tBalance: %s\tLimit: %s\n", customer.Name, customer.Balance, customer.CreditLimit)
	return details
}
