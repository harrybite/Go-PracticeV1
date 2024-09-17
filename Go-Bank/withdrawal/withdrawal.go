package withdrawal

// "fmt"
import (
	"fmt"
	"go-bank/usersHandler"
	"strconv"
)

func WithdrawalMoney() error {
	database, err := usersHandler.ConnectWithDB()
	if err != nil {
		return err
	}
	var address string
	var amount float64
	var pin int

	fmt.Println("Enter your address: ")
	fmt.Scan(&address)

	fmt.Println("Enter amount")
	fmt.Scan(&amount)

	fmt.Println("Enter pin")
	fmt.Scan(&pin)

	queryRow := `SELECT pin, balance FROM users WHERE address = $1`

	row := database.QueryRow(queryRow, address)

	var fetchedpin int64
	var fetchedbalance string
	err = row.Scan(&fetchedpin, &fetchedbalance)
	if err != nil {
		fmt.Println("error while fetching data: ", err)
		return err
	}
	if fetchedpin != int64(pin) {
		return fmt.Errorf("invalid pin")
	}
	balanceFloat, _ := strconv.ParseFloat(fetchedbalance, 64)
	if fetchedpin == int64(pin) {
		if balanceFloat < amount {
			return fmt.Errorf("insufficient fund")
		}
		updateSQL := `UPDATE users SET balance = balance - $1 WHERE address = $2`
		_, err = database.Exec(updateSQL, amount, address)
		if err != nil {
			return err
		}
		fmt.Println("Amount Deposited")
		fmt.Println("Left balance: ", balanceFloat-amount)
	}
	return nil
}
