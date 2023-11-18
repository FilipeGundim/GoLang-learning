package main

import "fmt"

func main() {

	var x int = 10
	var y int = x

	fmt.Println(x, y)

	x++
	// y doenst change
	fmt.Println(x, y)

	// pointers is a memory reference

	var var3 int = 100
	var pointer *int = &var3

	fmt.Println(var3, pointer)
	// 100 0xc00000a100

	fmt.Println(var3, *pointer)
	// 100 100

	var3 = 150
	fmt.Println(var3, *pointer)
}
