package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello, World!", channel)

	// for {
	// 	message, open := <-channel

	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(message)
	// }

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End program")
}

func write(text string, channel chan string) {
	for range 5 {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
