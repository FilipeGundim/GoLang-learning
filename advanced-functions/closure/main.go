package main

import "fmt"

func closure() func() {
	user := map[string]string{
		"name": "Filipe",
	}

	return func() {
		fmt.Println(user)
	}
}

func main() {
	fmt.Println("Inside main")

	newFunc := closure()
	newFunc()
}
