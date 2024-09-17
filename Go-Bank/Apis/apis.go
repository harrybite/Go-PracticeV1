package Apis

import (
	"encoding/json"
	"fmt"
	usersHandler "go-bank/usersHandler"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

type UserAddress struct {
	Address string `json:"useraddress"`
}

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var greet string = "hello from server"
	jsonUsers, err := json.Marshal(greet)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

func CheckUserBalance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var useraddress UserAddress
	err := json.NewDecoder(r.Body).Decode(&useraddress)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	exists, name, age, balance, err := usersHandler.UserChecker(useraddress.Address)

	if err != nil {
		http.Error(w, "Error while reading user", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"exists":  exists,
		"name":    name,
		"age":     age,
		"balance": balance,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	// Write the JSON response to the client
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body into User struct
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Add the new user to the users slice
	users = append(users, newUser)

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User added successfully")
}
