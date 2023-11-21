package main

import "fmt"

func invertSinal(number int) int {

	return number * -1
}

func invertPointerSinal(number *int) {
	*number = *number * -1
}

func main() {

	number := 20
	invertedNumder := invertSinal(number)

	fmt.Println(invertedNumder)

	newNumber := 40
	fmt.Println(newNumber)
	invertPointerSinal(&newNumber)
	fmt.Println(newNumber)
}
