package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Init main")

	err := checkmail.ValidateFormat("filipegundim@gmaill.com")
	fmt.Println(err)
}
