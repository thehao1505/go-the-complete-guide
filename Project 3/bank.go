package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERRORS")
		fmt.Println(err)
		fmt.Println("______________________")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to the bank!")
	fmt.Println("Contact us 24/7: ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount! Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 || withdrawalAmount > accountBalance {
				fmt.Println("Invalid withdrawal amount! Must be greater than 0 and less than your account balance.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your new balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}

// if choice == 1 {
// 	fmt.Println("Your balance is: ", accountBalance)
// } else if choice == 2 {
// 	var depositAmount float64
// 	fmt.Print("Enter the amount you want to deposit: ")
// 	fmt.Scan(&depositAmount)

// 	if depositAmount <= 0 {
// 		fmt.Println("Invalid deposit amount! Must be greater than 0.")
// 		// return
// 		continue
// 	}

// 	accountBalance += depositAmount
// 	fmt.Println("Your new balance is: ", accountBalance)
// } else if choice == 3 {
// 	var withdrawalAmount float64
// 	fmt.Print("Enter the amount you want to withdraw: ")
// 	fmt.Scan(&withdrawalAmount)

// 	if withdrawalAmount <= 0 || withdrawalAmount > accountBalance {
// 		fmt.Println("Invalid withdrawal amount! Must be greater than 0 and less than your account balance.")
// 		// return
// 		continue
// 	}

// 	accountBalance -= withdrawalAmount
// 	fmt.Println("Your new balance is: ", accountBalance)
// } else {
// 	fmt.Println("Goodbye!")
// 	// return
// 	break
// }
