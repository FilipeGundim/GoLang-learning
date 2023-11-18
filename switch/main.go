package main

import (
	"errors"
	"fmt"
)

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	case 8:
		fallthrough
	default:
		return errors.New("Invalid day of week switch between 1 -> 7").Error()
	}
}

func main() {
	fmt.Println("Switch")

	fmt.Println(dayOfWeek(1))
	fmt.Println(dayOfWeek(2))

	fmt.Println(dayOfWeek(8))
}
