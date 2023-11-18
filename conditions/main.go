package main

import "fmt"

func main() {

	fmt.Println("Conditions")

	number := 10

	if number > 15 {
		fmt.Println(number, "Is bigger than 15")
	} else {
		fmt.Println("Is less or equal than 15")
	}

	// ifinit

	if otherNumber := number; otherNumber > 0 {
		fmt.Println(otherNumber, "Is bigger than 0")
	} else {
		fmt.Println(otherNumber, "Is less or equal than 0")
	}

	// we cant access variables created with ifinit outsed the scope
	// undefined: otherNume
	// fmt.Println(otherNumer)
}
