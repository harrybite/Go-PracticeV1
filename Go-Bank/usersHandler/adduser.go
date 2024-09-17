package usersHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	// "net/http"
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

var DataBase *sql.DB

func AddNewUser() (string, error) {
	// bank, err := ReadBankData()
	err := CreatingBankDatabase()

	if err != nil {
		return "", fmt.Errorf("error reading bank data: %v", err)
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

	newUser := User{Name: name, Age: age, Balance: balance, Address: address, Pin: pin}
	errors := InsertUser(newUser)

	if errors != nil {
		return "", fmt.Errorf("error writing bank data: %v", err)
	}
	return fmt.Sprintln("User update successfully"), nil
}

func UserChecker(userAddress string) (bool, string, int, int64, error) {
	bank, err := ReadBankData()
	if err != nil {
		return false, "", 0, 0, fmt.Errorf("error reading bank data: %v", err)

	}
	var address string
	if userAddress == "" {
		fmt.Println("Enter your address:")
		if _, err := fmt.Scan(&address); err != nil {
			return false, "", 0, 0, fmt.Errorf("error reading address: %v", err)
		}
	} else {
		address = userAddress
	}

	if value, exists := bank.Users[address]; exists {
		return exists, value.Name, value.Age, value.Balance, nil
		// fmt.Sprint("\nExist %v\nName %v\nAge %v\nBalance $%v\n\n", exists, value.Name, value.Age, value.Balance)
	} else {
		return false, "", 0, 0, fmt.Errorf("User not found")
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

func CreatingBankDatabase() error {
	database, err := ConnectWithDB()

	if err != nil {
		fmt.Println("Error while connecting with database: ", err)
		return err
	}

	err = database.Ping()
	if err != nil {
		fmt.Println("database connection error: ", err)
		return err
	}

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			address TEXT PRIMARY KEY,
			name TEXT,
			age INT,
			pin INT,
			balance DECIMAL(10, 2)
		);
	`
	_, err = database.Exec(createTableSQL)
	if err != nil {
		log.Fatal("error creating table: ", err)
		return err
	}
	fmt.Println("users Table created successfully.")
	return nil
}

func InsertUser(user User) error {
	insertSQL := `
		INSERT INTO users (name, age, pin, balance, address)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := DataBase.Exec(insertSQL, user.Name, user.Age, user.Pin, user.Balance, user.Address)
	if err != nil {
		fmt.Println("error while inserting into table: ", err)
		return err
	} else {
		fmt.Println("from InsertUser: inserted data successfully")
	}
	return err
}

func ConnectWithDB() (*sql.DB, error) {
	var dbstr string = "postgres://nikhil:nikhil@localhost/nikhil?sslmode=disable"
	database, err := sql.Open("postgres", dbstr)
	DataBase = database
	if err != nil {
		return nil, err
	}
	return database, nil
}

func ReadData(address string) error {
	database, err := ConnectWithDB()
	if err != nil {
		fmt.Println("error connecting with database: ", err)
		return err
	}
	defer database.Close()

	querySQL := `
	SELECT name, age, balance, address FROM users WHERE address = $1
	`

	row := database.QueryRow(querySQL, address)
	var name string
	var age int
	var balance float64
	var fetchedAddress string

	err = row.Scan(&name, &age, &balance, &fetchedAddress)
	if err != nil {
		fmt.Println("error while fetching data: ", err)
		return err
	}
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Balance:", balance)
	fmt.Println("Address:", fetchedAddress)
	return nil
}
