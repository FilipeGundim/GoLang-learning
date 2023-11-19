package main

import "fmt"

func fibonacci(position uint) uint {

	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)

}

func main() {
	fmt.Println("Recursive functions")

	position := uint(10)
	fmt.Println(fibonacci(position))
}
