package main

import "fmt"

type person struct {
	name   string
	age    uint8
	height float32
}

type student struct {
	person
	course string
}

func main() {
	fmt.Println("Extends")

	p1 := person{"Filipe", 26, 1.82}
	fmt.Println(p1)

	e1 := student{p1, "ads"}
	fmt.Println(e1)
	fmt.Println(e1.name)
}
