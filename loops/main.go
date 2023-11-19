package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 20 {
	// 	i++
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i, "Icrement i")
	// }

	// for j := 0; j < 10; j++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(j, "Icrement j")
	// }

	numbers := [3]int{1, 2, 3}

	// I put _ to ignore first param (index of value)
	for _, number := range numbers {
		fmt.Println(number)
	}

	const name string = "Filipe"

	for idx, letter := range name {
		fmt.Println(idx, letter)
		fmt.Println(idx, string(letter))
	}

	user := map[string]string{
		"name": name,
	}

	fmt.Println(user)
	for key, value := range user {
		fmt.Println(key, value)
	}

	// for {
	// 	fmt.Println("This is a infinit loop")
	// 	time.Sleep(time.Second)
	// }

}
