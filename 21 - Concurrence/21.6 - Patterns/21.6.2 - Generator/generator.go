package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello, World!")

	for range 10 {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Got value: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
