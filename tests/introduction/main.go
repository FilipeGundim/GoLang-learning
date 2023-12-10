package main

import (
	"fmt"
	"tests/introduction/adresses"
)

func main() {
	typeAddress := adresses.TypeOfAddress("Avenida Paulista")
	fmt.Println(typeAddress)
}
