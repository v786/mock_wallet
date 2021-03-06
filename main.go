package main

import (
	"bufio"
	"customer"
	"fmt"
	"merchant"
	"os"
	"strconv"
	"strings"
	"transaction"
)

const (
	menu = `
------------Welcome--------------
	Main Menu:
	1 Add Customer: Name Amount
	2 Add Merchant
	3 List all Merchants
	4 Update balance - Serial-number Amount
	5 Print Menu
	6 List all customers with balance
	7 Update Merchant Discount
	8 Settle Merchant Balance
	0 Exit
`
)

var customerList []*customer.Customer
var merchantList []*merchant.Merchant
var transactionList []*transaction.Transaction

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(menu)
	for {
		fmt.Print("=> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "x":
		test()
	case "exit":
	case "0":
		os.Exit(0)
		// add another case here for custom commands.
	case "1":
		amt, err := strconv.Atoi(arrCommandStr[2])
		if err != nil {
			return err
		}
		addCustomer(arrCommandStr[1], amt)
		return nil
	case "2":
		getCustomers()
	case "3":
		getMerchants()
	case "4":
		updateBalance(arrCommandStr[1], arrCommandStr[2])
	case "5":
		fmt.Println(menu)
	case "6":
		disc, err := strconv.Atoi(arrCommandStr[2])
		if err != nil {
			return err
		}
		addMerchant(arrCommandStr[1], disc)
	case "7":

	}

	return nil
}

func getCustomers() {
	fmt.Println("Customer details : \nSlNo\tName\tBalance\tLimit")
	for i := range customerList {
		fmt.Println(i+1, "\t", customerList[i].GetDetails())
	}
}

func getMerchants() {
	fmt.Println("Merchant details : \nSlNo\tName\tBalance\tDiscount")
	for i := range merchantList {
		fmt.Println(i+1, "\t", merchantList[i].GetDetails())
	}
}

func getTransactions() {
	fmt.Println("Transaction details : \nSlNo\tMerch\tCust\tAmount\tStatus")
	for i := range transactionList {
		fmt.Println(i+1, "\t", transactionList[i].GetDetails())
	}
}

func addCustomer(name string, amt int) {
	fmt.Println("Creating customer ....")
	var customerNew customer.Customer
	customerNew.Create(name, amt)
	customerList = append(customerList, &customerNew)
}

func updateBalance(custIndex string, amount string) error {

	fmt.Println("Updating customer ....")

	amt, err := strconv.Atoi(amount)

	if err != nil {
		return err
	}

	index, err := strconv.Atoi(custIndex)

	if err != nil {
		return err
	}

	cust := customerList[index]

	return cust.UpdateBalance(amt)
}

func addMerchant(name string, disc int) {
	fmt.Println("Creating merchant ....")
	var newMerchant merchant.Merchant
	newMerchant.Create(name, disc)
	merchantList = append(merchantList, &newMerchant)
}

func makeTransaction(custIndex string, merchIndex string, amt string) error {

	index, err := strconv.Atoi(custIndex)
	if err != nil {
		return err
	}

	cust := customerList[index]

	index, err = strconv.Atoi(merchIndex)
	if err != nil {
		return err
	}

	merchant := merchantList[index]

	amount, err := strconv.Atoi(amt)
	if err != nil {
		return err
	}

	var t transaction.Transaction
	transactionList = append(transactionList, &t)

	err = t.Create(cust, merchant, amount)
	if err != nil {
		return err
	}

	return nil
}

func getTransactionsByCustomer(cust customer.Customer) []transaction.Transaction {

	var tranList []transaction.Transaction
	var t transaction.Transaction

	for i := range transactionList {
		t = *transactionList[i]
		if cust == *t.Customer {
			tranList = append(tranList, t)
		}
	}

	return tranList
}

func printTransList(tranList []transaction.Transaction) {
	fmt.Println("Transaction details by Customer : \nSlNo\tMerch\tCust\tAmount\tStatus")
	for i := range tranList {
		fmt.Println(i+1, "\t", transactionList[i].GetDetails())
	}
}

func getCustomerByIndex(index int) customer.Customer {

	var cust customer.Customer

	cust = *customerList[index]

	return cust
}

func test() {
	addCustomer("ABC", 100)
	addCustomer("AB1", 200)
	addCustomer("AB2", 300)
	addCustomer("AB3", 400)
	addMerchant("mBC", 10)
	addMerchant("mB1", 20)
	addMerchant("mB2", 30)
	addMerchant("mB3", 4)

	updateBalance("1", "300")

	fmt.Println(makeTransaction("1", "1", "50"))

	c := getCustomerByIndex(1)
	tl := getTransactionsByCustomer(c)

	printTransList(tl)
	getTransactions()
	getCustomers()
	getMerchants()
	os.Exit(0)
}
