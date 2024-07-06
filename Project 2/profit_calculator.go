package main

import (
	"fmt"
	"os"
)

const resultFile = "result.txt"

func writeResultToFile(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile(resultFile, []byte(result), 0644)
}

func main() {
	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		fmt.Println("Input can not negative or equal zero.")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// writeResultToFile(ebt, profit, ratio)
	writeResultToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
