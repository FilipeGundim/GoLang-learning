package main

import (
	"errors"
	"fmt"
)

func main() {
	//Quantity of bits
	//int, int8, int16, int32, int64

	//unsyged int
	//uint

	var number int16 = 100
	fmt.Println(number)

	var number2 uint = 10000
	fmt.Println(number2)

	// alias
	// INT32 = RUNE
	var number3 rune = 123456
	fmt.Println(number3)

	//byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

	var realNumber float32 = 123.45
	fmt.Println(realNumber)

	var realNumber2 float64 = 120000000000.45
	fmt.Println(realNumber2)

	realNumber3 := 12345.67
	fmt.Println(realNumber3)

	str := "B"
	character := 'B'
	fmt.Println(str, character)

	var boolean1 bool = true
	var boolean2 bool
	fmt.Println(boolean1, boolean2)

	var err1 error
	fmt.Println(err1)

	var err2 error = errors.New("undefined")
	fmt.Println(err2)
}
