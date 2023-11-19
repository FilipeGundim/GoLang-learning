package main

import "fmt"

func recoverExecution() {
	fmt.Println("Recover function")

	if r := recover(); r != nil {
		fmt.Println("Recover execution")
	}
}

func studentIsApproved(n1 float64, n2 float64) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("Average is 6")
}

func main() {
	fmt.Println(studentIsApproved(6, 7))
	fmt.Println("After execution")

	fmt.Println(studentIsApproved(6, 6))
}
