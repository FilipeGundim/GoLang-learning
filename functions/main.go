package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	soma := somar(8, 8)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("frutenfraten")
	fmt.Println(result)

	sum, sub := mathCalcs(10, 15)
	fmt.Println(sum, sub)

	// ignoring a return param
	sum2, _ := mathCalcs(10, 15)
	fmt.Println(sum2)
}
