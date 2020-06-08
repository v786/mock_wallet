package merchant

import "fmt"

//Merchant : Users that accept wallet balance in return of something
type Merchant struct {
	Name     string
	Discount int
	Balance  int
	Sale     int
}

func (merchant *Merchant) Create(name string, discount int) error {
	merchant.Name = name
	merchant.Discount = discount

	fmt.Println("merchant Created:", merchant.Name, merchant.Discount)

	return nil
}

func (merchant *Merchant) GetDetails() string {
	details := fmt.Sprintf("%s\t%d\t%d\n", merchant.Name, merchant.Balance, merchant.Discount)
	return details
}
