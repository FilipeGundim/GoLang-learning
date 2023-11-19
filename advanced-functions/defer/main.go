package main

import "fmt"

func function1() {
	fmt.Println("Executing function1")
}

func function2() {
	fmt.Println("Executing function2")
}

func function3() {
	fmt.Println("Executing function3")
}

func main() {
	defer function1()
	// defer makes Go execute the function is the last possible moment
	defer function3()
	function2()
	function2()
}
