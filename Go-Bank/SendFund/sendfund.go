package sendfund

import (
	"fmt"
	"go-bank/usersHandler"
)

func SendMoney() error {
	var senderaddress string
	var receiveraddress string
	var amount int64
	var pin int

	bank, err := usersHandler.ReadBankData()
	if err != nil {
		return fmt.Errorf("error reading bank data: %v", err)
	}

	fmt.Println("Enter your address:")
	fmt.Scan(&senderaddress)

	if _, exists := bank.Users[senderaddress]; !exists {
		fmt.Println("User not found")
		return nil
	}

	fmt.Println("Enter your pin:")
	fmt.Scan(&pin)

	// Check if the entered PIN is correct
	if bank.Users[senderaddress].Pin != pin {
		fmt.Println("Incorrect PIN")
		return nil
	}

	fmt.Println("Enter amount:")
	fmt.Scan(&amount)

	// Validate the withdrawal amount
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return nil
	}

	// Check if the withdrawal amount is greater than the user's balance
	if amount > bank.Users[senderaddress].Balance {
		fmt.Println("Insufficient balance")
		return nil
	}

	fmt.Println("Enter receiver address:")
	fmt.Scan(&receiveraddress)

	if _, exists := bank.Users[receiveraddress]; !exists {
		fmt.Println("receiver not found")
		return nil
	}

	bank.Users[senderaddress].Balance -= amount
	bank.Users[receiveraddress].Balance += amount

	// Write updated bank data back to the file
	err = usersHandler.WriteBankData(bank)
	if err != nil {
		return fmt.Errorf("error writing bank data: %v", err)
	}

	fmt.Println("Money transfer successfully")
	return nil

}
