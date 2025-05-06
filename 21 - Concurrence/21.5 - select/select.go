package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		select {
		case channelMessage1 := <-channel1:
			fmt.Println(channelMessage1)
		case channelMessage2 := <-channel2:
			fmt.Println(channelMessage2)
		}
	}
}
