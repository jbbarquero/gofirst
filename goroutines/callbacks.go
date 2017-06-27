package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main() {

	fmt.Println("Sync callback")

	toUpperSync("Hello Callbacks", func(value string) {
		fmt.Printf("Sync Callback: %s\n", value)
	})

	fmt.Println("Async callback")

	wait.Add(1)
	toUpperAsync("Hello Callbacks", func(value string) {
		fmt.Printf("Async Callback: %s\n", value)
		wait.Done()
	})
	fmt.Println("Waiting async response...")
	wait.Wait()

	fmt.Println("End callbacks")
}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
