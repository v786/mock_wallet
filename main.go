package main

import (
	"bufio"
	"customer"
	"fmt"
	"os"
	"strings"
)

var customerList []customer.Customer

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`------------Welcome--------------
		Main Menu:
		1 Name Amount
		2 List all customers with balance
		3 List all Merchants
		4 Update balance - Serial-number Amount
		5 Print Menu`)
	for {
		fmt.Print("$ ")
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
	case "exit":
		os.Exit(0)
		// add another case here for custom commands.
	case "1":
		addCustomer(arrCommandStr[1])
		return nil
	case "2":
		getCustomers()
	}

	return nil
}

func getCustomers() {
	fmt.Println("Customer details : SlNo-Name-Balance-Limit")
	for i := range customerList {
		fmt.Println(i, customerList[i].GetDetails())
	}
}

func addCustomer(name string) {
	fmt.Println("Creating customer ....")
	var customerNew customer.Customer
	customerNew.Create(name, 100)
	customerList = append(customerList, customerNew)
}
