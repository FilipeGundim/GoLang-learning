package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello world")
	write("Hello world with Go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
