package withdrawal

import (
	"fmt"
	users "go-bank/usersHandler"
)

func WithdrawalMoney() error {
	var address string
	var amount int64
	var pin int
	bank, err := users.ReadBankData()
	if err != nil {
		return fmt.Errorf("error reading bank data: %v", err)
	}

	fmt.Println("Enter your address:")
	fmt.Scan(&address)

	// Check if the user exists
	if _, exists := bank.Users[address]; !exists {
		fmt.Println("User not found")
		return nil
	}

	fmt.Println("Enter your pin:")
	fmt.Scan(&pin)

	// Check if the entered PIN is correct
	if bank.Users[address].Pin != pin {
		fmt.Println("Incorrect PIN")
		return nil
	}

	fmt.Println("Enter amount to withdraw:")
	fmt.Scan(&amount)

	// Validate the withdrawal amount
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return nil
	}

	// Check if the withdrawal amount is greater than the user's balance
	if amount > bank.Users[address].Balance {
		fmt.Println("Insufficient balance")
		return nil
	}

	// Update user's balance
	bank.Users[address].Balance -= amount

	// Write updated bank data back to the file
	err = users.WriteBankData(bank)
	if err != nil {
		return fmt.Errorf("error writing bank data: %v", err)
	}

	fmt.Println("Withdrawal successful")
	return nil
}
