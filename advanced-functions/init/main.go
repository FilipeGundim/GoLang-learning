package main

import "fmt"

func init() {
	fmt.Println("Executing init -> exec before main function")
}

func main() {
	fmt.Println("Executing main")
}
