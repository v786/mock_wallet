package transaction

import (
	"customer"
	"errors"
	"fmt"
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

	transaction.Customer = cust
	transaction.Merchant = merch
	transaction.Amount = amount
	transaction.Discount = merch.Discount

	var cost int
	if amount > merch.Discount {
		cost = amount - merch.Discount
	} else {
		cost = 0
	}

	transaction.Success = false

	fmt.Println("Begin transaction....", cost, amount, cust.Balance)

	if cust.Balance >= amount {
		cust.Balance = cust.Balance - amount
		merch.Balance = merch.Balance + cost
		transaction.Success = true
	} else {
		return errors.New("balance not sufficient")
	}

	return nil
}
