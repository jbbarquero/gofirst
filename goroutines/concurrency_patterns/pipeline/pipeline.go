package pipeline

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n := os.Args[0] //https://gobyexample.com/command-line-arguments
	i, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("Error with the number provided %v. Usage pipeline number.")
	}
	//Instead of https://gobyexample.com/command-line-flags

	go LaunchPipeline(i)
}

func LaunchPipeline(amount int) int {
	channelElements := generator(amount)
	channelSquares := power(amount, channelElements)
	result := addUpAll(channelSquares)

	sum := <-result

	fmt.Printf("Pipeline result: %d\n", sum)

	return sum
}

func generator(count int) <-chan int {
	channelElement := make(chan int)
	go func() {
		for i := 0; i <= count; i++ {
			channelElement <- i
		}
		close(channelElement)
	}()
	return channelElement
}

func power(count int, in <-chan int) <-chan int {
	out := make(chan int, count)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func addUpAll(channelNumbers <-chan int) chan int {
	channelResult := make(chan int)
	sum := 0
	go func() {
		for number := range channelNumbers {
			sum += number
		}
		channelResult <- sum
		close(channelResult)
	}()
	return channelResult
}
