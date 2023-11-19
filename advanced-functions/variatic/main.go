package main

import (
	"errors"
	"fmt"
)

func sum(numbers ...int) (total int, err error) {
	if len(numbers) == 0 {
		err = errors.New("No number was received")
	}

	for _, number := range numbers {
		total += number
	}

	return

	// implicit return total, err
}

func main() {
	total, err := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(total)
	}

	total2, err2 := sum()

	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println(total2)
	}
}
