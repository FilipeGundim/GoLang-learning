package main

import "fmt"

func main() {
	sum := 1 + 2
	sub := 1 - 2
	mult := 1 * 2
	div := 1 / 2
	rest := 10 % 2

	fmt.Println(sum, sub, mult, div, rest)

	fmt.Println("---------------------")

	// var number1 int16 = 10
	// var number2 int32 = 25

	// sum1 := number1 + number2;
	// error

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println("---------------------")

	fmt.Println(true && true)
	tr, fl := true, false
	fmt.Println(tr && fl)
	fmt.Println(tr || fl)
	fmt.Println(!fl)

	fmt.Println("---------------------")

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	numero += 3
	numero /= 10
	fmt.Println(numero)

	fmt.Println("---------------------")

	// texto := numero > 5 ? "bigger than five"  : "less than five"
	var texto string
	if numero > 5 {
		texto = "bigger than five"
	} else {
		texto = "less than five"
	}
	fmt.Println(texto)
}
