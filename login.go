package main

import (
	"fmt"
)

func main() {
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	// Predefined credentials
	const trueUsername = "user"
	const truePassword = "pass"

	// Check credentials
	if username == trueUsername && password == truePassword {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed. Incorrect username or password.")
	}
}