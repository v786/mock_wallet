package user

import "fmt"

//User : End user that will use the wallet
type User struct {
	Name        string
	Balance     int
	CreditLimit int
}

func (user *User) Create(name string, creditLimit int) error {
	user.Name = name
	user.CreditLimit = creditLimit
	user.Balance = creditLimit

	fmt.Println("User Created:", user.Name, user.CreditLimit)

	return nil
}

func (user *User) GetDetails() error {
	fmt.Println("Filed:", "Details")

	fmt.Println("Name:", user.Name)
	fmt.Println("Balance:", user.Balance)
	fmt.Println("Limit:", user.CreditLimit)

	return nil
}
