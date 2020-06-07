package main

import (
	"bufio"
	"fmt"
	user "mock_wallet/user"
	"os"
	"strings"
)

var userList []user.User

func main() {
	reader := bufio.NewReader(os.Stdin)
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
	case "user":
		fmt.Println("Creating User")
		var userNew user.User
		userNew.Create(arrCommandStr[1], 100)
		userNew.GetDetails()
	}

	return nil
}
