package main

import "fmt"

func mathCalcs(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2

	return

	// implicit return sum, sub
}

func main() {

	fmt.Println("Named return")

	sum, sub := mathCalcs(10, 20)
	fmt.Println(sum, sub)
}
