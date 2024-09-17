package deposit

import (
	"fmt"
	"go-bank/usersHandler"
)

func DepositAmout() error {
	database, err := usersHandler.ConnectWithDB()
	if err != nil {
		return err
	}
	defer database.Close()
	var address string
	var amount int64

	fmt.Println("Enter your address: ")
	fmt.Scan(&address)

	fmt.Println("Enter amount")
	fmt.Scan(&amount)

	updateSQL := `UPDATE users SET balance = balance + $1 WHERE address = $2`
	_, err = database.Exec(updateSQL, amount, address)
	if err != nil {
		return err
	}
	fmt.Println("Balance updated successfully")
	return nil
}
