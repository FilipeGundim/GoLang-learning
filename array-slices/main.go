package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array & Slices")

	var arr1 [5]int
	arr1[0] = 1

	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr3)

	// arr3[8] = 8
	// error

	slice := []int{13, 14, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))

	fmt.Println(slice)
	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)

	fmt.Println(slice2, arr2)
	arr2[2] = 5
	fmt.Println(slice2, arr2)

	// intern arrays
	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 5)
	slice4 = append(slice4, 5)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
