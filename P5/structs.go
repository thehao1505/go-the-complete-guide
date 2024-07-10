package main

import (
	"fmt"
)

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthdate (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
