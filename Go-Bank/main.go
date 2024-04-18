package main

import (
	"fmt"
	sendfund "go-bank/SendFund"
	deposit "go-bank/deposit"
	users "go-bank/usersHandler"
	withdraw "go-bank/withdrawal"
)

func main() {
	var choise int
	fmt.Println("Welcome to Go-Bank!")
	fmt.Println("Select the following operation")

	fmt.Println("1. Check User")
	fmt.Println("2. Add user")
	fmt.Println("3. Deposit Money")
	fmt.Println("4. Withdrwal Money")
	fmt.Println("5. Send Money")

	fmt.Scan(&choise)

	if choise == 1 {
		users.UserChecker()
	} else if choise == 2 {
		users.AddNewUser()
	} else if choise == 3 {
		deposit.DepositAmout()
	} else if choise == 4 {
		withdraw.WithdrawalMoney()
	} else if choise == 5 {
		sendfund.SendMoney()
	} else {
		fmt.Println("Thank you for visiting us!")
	}

}
