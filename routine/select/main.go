package main

import (
	"fmt"
	"time"
)

const FIVE_HUNDRED_MILLISECONDS, TWO_SECONDS = time.Millisecond * 500, time.Second * 2

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(FIVE_HUNDRED_MILLISECONDS)
			channel1 <- "Channel 1"
		}

	}()

	go func() {
		for {
			time.Sleep(TWO_SECONDS)
			channel2 <- "Channel 2"
		}
	}()

	for {

		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}
}
