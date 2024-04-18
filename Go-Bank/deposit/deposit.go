package deposit

import (
	"fmt"
	users "go-bank/usersHandler"
)

func DepositAmout() error {
	var address string
	var amount int64
	bank, _ := users.ReadBankData()

	fmt.Println("Enter your address")
	fmt.Scan(&address)

	if _, exists := bank.Users[address]; !exists {
		fmt.Printf("User could not find")
		return nil
	}

	fmt.Println("Enter amount")
	fmt.Scan(&amount)

	// Update user balance
	bank.Users[address].Balance += amount

	// Write updated bank data back to the file
	err := users.WriteBankData(bank)
	if err != nil {
		return fmt.Errorf("error writing bank data: %v", err)
	}
	fmt.Printf("\nUser balance updated successfully\n\n")
	return nil
}
