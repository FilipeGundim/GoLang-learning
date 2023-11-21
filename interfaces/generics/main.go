package main

import (
	"fmt"
	"reflect"
)

func generics(interf interface{}) {
	fmt.Println(interf)
	fmt.Println(reflect.TypeOf(interf))
}

func main() {
	slice := []int{2, 2, 2}
	generics(10)
	generics("10")
	generics(2.1)
	generics(slice)
}
