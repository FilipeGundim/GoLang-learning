package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplex(write("Hello world"), write("Hello world2"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(channelIn1, channelIn2 <-chan string) <-chan string {
	channelOut := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelIn1:
				channelOut <- message

			case message := <-channelIn2:
				channelOut <- message
			}
		}
	}()

	return channelOut
}

func write(text string) <-chan string {

	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
