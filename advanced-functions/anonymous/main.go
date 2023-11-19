package main

import "fmt"

func main() {

	received := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Hello anonymous world")

	fmt.Println(received)
}
