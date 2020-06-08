package transaction

import (
	"customer"
	"errors"
	"merchant"
)

//Transaction : Entity that stores a transaction details between customer and merchant
type Transaction struct {
	Customer *customer.Customer
	Merchant *merchant.Merchant
	Amount   int
	Discount int
	Success  bool
}

func (transaction *Transaction) Create(cust *customer.Customer, merch *merchant.Merchant, amount int) error {

	if cust.Balance >= amount {
		cust.Balance = cust.Balance - amount
	} else {
		return errors.New("balance not sufficient")
	}

	return nil
}
