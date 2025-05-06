package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Hello, World!")
		waitGroup.Done()
	}()

	go func() {
		write("Programming in Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for range 5 {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
