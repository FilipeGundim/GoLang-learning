package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello world", channel)

	mensage, isOpen := <-channel

	if isOpen {
		fmt.Println("channel is open")
	}

	fmt.Println(mensage)
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
