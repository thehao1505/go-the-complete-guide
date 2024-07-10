package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	getAdultYears(&age)
	// adultYears := getAdultYears(agePointer)
	// fmt.Println("Adult Years: ", adultYears)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
