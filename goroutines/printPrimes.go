package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	numProcessors := runtime.NumCPU()
	fmt.Printf("Init for %d processors\n", numProcessors)
	runtime.GOMAXPROCS(numProcessors)
}

func main() {
	fmt.Println("Main Print primes...")

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go PrintPrimes("A")
	go PrintPrimes("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Main Print primes. End.")
}

func PrintPrimes(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 10000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)

	fmt.Printf("Print primes %s.\n", prefix)

}
