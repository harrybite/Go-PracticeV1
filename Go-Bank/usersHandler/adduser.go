package usersHandler

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Balance int64  `json:"balance"`
	Address string `json:"address"`
	Pin     int    `json:"pin"`
}

type Bank struct {
	Name  string           `json:"name"`
	Users map[string]*User `json:"users"`
}

func AddNewUser() error {
	bank, err := ReadBankData()
	if err != nil {
		return fmt.Errorf("error reading bank data: %v", err)
	}
	var name string
	var age int
	var balance int64
	var address string
	var pin int

	fmt.Println("Enter you name")
	fmt.Scan(&name)
	fmt.Println("Enter you age")
	fmt.Scan(&age)
	fmt.Println("Enter you address")
	fmt.Scan(&address)
	fmt.Println("Enter you pin")
	fmt.Scan(&pin)
	fmt.Println("Enter deposit amount")
	fmt.Scan(&balance)

	newUser := &User{Name: name, Age: age, Balance: balance, Address: address, Pin: pin}
	bank.Users[address] = newUser

	err = WriteBankData(bank)
	if err != nil {
		return fmt.Errorf("error writing bank data: %v", err)
	}

	fmt.Println("User added successfully.")
	return nil

}

// func (u *User) String() string {
//     return fmt.Sprintf("Name: %s\nAge: %d\nBalance: %d\nAddress: %s\nPin: %d",
//         u.Name, u.Age, u.Balance, u.Address, u.Pin)
// }

func UserChecker() {
	var address string = "0x3F382Db2D9B9AeD2570c296Faa71e98e90afD352"
	// var pin int

	fmt.Println("Enter you address")
	fmt.Scan(&address)
	// fmt.Println("Enter you pin")
	// fmt.Scan(&pin)

	bank, _ := ReadBankData()

	if value, exists := bank.Users[address]; exists {
		fmt.Printf("\nExist %v\nName %v\nAge %v\nBalance $%v\n\n", exists, value.Name, value.Age, value.Balance)
	}
}

func ReadBankData() (*Bank, error) {
	file, err := os.Open("/Users/nikhil/Downloads/Github/Go-PracticeV1/Go-Bank/usersHandler/Data/users.json")
	if err != nil {
		fmt.Printf("error opening bank data: %v", err)
		return nil, err
	}
	// closing the file from memory
	defer file.Close()

	var bank Bank
	err = json.NewDecoder(file).Decode(&bank)
	if err != nil {
		fmt.Printf("error reading bank data: %v", err)
		return nil, err
	}
	return &bank, nil
}

func WriteBankData(bank *Bank) error {
	file, err := os.Create("/Users/nikhil/Downloads/Github/Go-PracticeV1/Go-Bank/usersHandler/Data/users.json")
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(bank)
	if err != nil {
		return err
	}
	return nil
}
