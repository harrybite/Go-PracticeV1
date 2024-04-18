package main

import (
	"fmt"
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

	fmt.Scan(&choise)

	if choise == 1 {
		users.UserChecker()
	} else if choise == 2 {
		users.AddNewUser()
	} else if choise == 3 {
		deposit.DepositAmout()
	} else if choise == 4 {
		withdraw.WithdrwalMoney()
	} else {
		fmt.Println("Thank you for visiting us!")
	}

}

func WithDrwalMoney(balance float64) float64 {
	var value float64
	fmt.Print("Enter amount to withdrwal: ")
	fmt.Scan(&value)
	if value < balance {
		balance = balance - value
	} else {
		fmt.Println("insufficient amount")
	}
	return balance
}
