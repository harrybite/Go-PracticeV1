package userHandler

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

func adduser() error {
	bank, err := ReadBankData("/Users/nikhil/Downloads/Github/GoPractice/Go-Bank/Data/users.json")
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
	bank.Users[name] = newUser

	err = WriteBankData("/Users/nikhil/Downloads/Github/GoPractice/Go-Bank/Data/users.json", bank)
	if err != nil {
		return fmt.Errorf("error writing bank data: %v", err)
	}

	fmt.Println("User added successfully.")
	return nil

}

func deleteuser() {

}

func ReadBankData(filename string) (*Bank, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// closing the file from memory
	defer file.Close()

	//
	var bank Bank
	err = json.NewDecoder(file).Decode(&bank)
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func WriteBankData(filename string, bank *Bank) error {
	file, err := os.Create(filename)
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
