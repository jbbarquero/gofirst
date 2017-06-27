package main

import "time"

func main() {
	go helloWorld()

	go func() {
		println("Hello Inner World!")
	}()

	go func(msg string) {
		println(msg)
	}("Hello Inner World with parameters!")

	messagePrinter := func(msg string) {
		println(msg)
	}

	go messagePrinter("Hello variable World!")
	go messagePrinter("Hello variable World one more time!")

	time.Sleep(time.Second)

}

func helloWorld() {
	println("Hello World!")
}
