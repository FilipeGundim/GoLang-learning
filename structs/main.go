package main

import "fmt"

type user struct {
	name    string
	age     int
	address address
}

type ar struct {
	number int
	name   string
}

func main() {
	fmt.Println("Structs")

	var u user
	fmt.Println(u)

	u.name = "Filipe"
	u.age = 25
	fmt.Println(u)

	u2 := user{"Patricia", 3, address{30, "Georgina sa leite"}}
	fmt.Println(u2)
}
