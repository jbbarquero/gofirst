package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hello World!"
	}()

	message := <-channel
	fmt.Println(message)

	channel2 := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		channel2 <- "Bye, bye, miss American Pie"
		fmt.Println("Finishing goroutine")
		waitGroup.Done()
	}()
	time.Sleep(time.Second)

	message2 := <-channel2
	fmt.Println(message2)
	waitGroup.Wait()
}
