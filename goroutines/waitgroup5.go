package main

import (
	"fmt"
	"sync"
)

func main() {

  var wait sync.WaitGroup
	var goRoutines = 5
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(gorutineId int) {
			fmt.Printf("ID:%d: goroutine\n", gorutineId)
			wait.Done()
		}(i)
	}

	wait.Wait()

}
