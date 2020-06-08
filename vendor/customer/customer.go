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
	details := fmt.Sprintf("%s\t%d\t%d\n", customer.Name, customer.Balance, customer.CreditLimit)
	return details
}

func (customer *Customer) UpdateBalance(amount int) error {
	customer.Balance = amount
	return nil
}
