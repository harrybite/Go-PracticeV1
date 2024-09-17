package main

import (
	"fmt"
	// apis "go-bank/Apis"

	"go-bank/deposit"
	users "go-bank/usersHandler"
	"go-bank/withdrawal"
	// "net/http"
)

func main() {

	// go func() {
	// 	http.HandleFunc("/", apis.GetHomeHandler)
	// 	http.HandleFunc("/users", apis.GetUsersHandler)
	// 	http.HandleFunc("/users/add", apis.AddUserHandler)
	// 	http.HandleFunc("/users/balance", apis.CheckUserBalance)

	// 	fmt.Println("Server listening on port 8080")
	// 	http.ListenAndServe(":8080", nil)
	// }()

	var choise int
	fmt.Println("Welcome to Go-Bank!")
	fmt.Println("Select the following operation")

	fmt.Println("1. Check User")
	fmt.Println("2. Add user")
	fmt.Println("3. Deposit Money")
	fmt.Println("4. Withdrwal Money")
	fmt.Println("5. Send Money")
	fmt.Println("6. Exit")

	fmt.Scan(&choise)

	switch choise {
	case 1:
		var address string
		fmt.Println("Enter user addresss")
		fmt.Scan(&address)
		users.ReadData(address)
	case 2:
		data, err := users.AddNewUser()
		if err != nil {
			fmt.Println("main.go: ", err)
		}
		fmt.Println(data)
	case 3:
		err := deposit.DepositAmout()
		if err != nil {
			fmt.Println("main.go: ", err)
		}
	case 4:
		err := withdrawal.WithdrawalMoney()
		if err != nil {
			fmt.Println("main.go: ", err)
		}
	case 6:
		break
	default:
		fmt.Println("Unknown")
		return
	}

	// if choise == 1 {
	// 	users.UserChecker("")
	// } else if choise == 2 {
	// 	users.AddNewUser()
	// } else if choise == 3 {
	// 	deposit.DepositAmout()
	// } else if choise == 4 {
	// 	withdraw.WithdrawalMoney()
	// } else if choise == 5 {
	// 	sendfund.SendMoney()
	// } else if choise == 6 {
	// 	// users.WriteBankDataP("postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	// 	users.CreatingBankDatabase()
	// } else {
	// 	fmt.Println("Thank you for visiting us!")
	// }

}
