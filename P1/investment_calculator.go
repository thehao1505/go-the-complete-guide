package main

import (
	"fmt"
	"math"
)

const inflationRate = 3.0

func main() {
	var investmentAmount, years, expectedReturn float64 // Automatically initialized to 0.0

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return: ")
	fmt.Scan(&expectedReturn)

	outputText("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturn, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("%.2f", futureValue)
	formattedFRV := fmt.Sprintf("%.2f", futureRealValue)
	fmt.Println(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturn, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturn/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
	// return
}
