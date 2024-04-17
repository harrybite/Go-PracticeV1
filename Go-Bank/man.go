package main

import (
	"fmt"
)

func main() {
	var balance float64 = 1000
	var choise int
	fmt.Println("Welcome to Go-Bank!")
	fmt.Println("Select the following operation")

	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdrwal Money")
	fmt.Printf("4.  Add user")

	fmt.Scan(&choise)

	if choise == 1 {
		printBalance(balance)
	} else if choise == 2 {
		balance := depositMoney(balance)
		printBalance(balance)
	} else if choise == 3 {
		balance := WithDrwalMoney(balance)
		printBalance(balance)
	} else if choise == 4 {

	} else {
		fmt.Println("Thank you for visiting us!")
	}

}

func depositMoney(balance float64) float64 {
	var value float64
	fmt.Print("Enter amount to deposite: ")
	fmt.Scan(&value)
	balance = balance + value
	return balance
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

func printBalance(balance float64) {
	fmt.Printf("Your balance: $%v\n", balance)
}
