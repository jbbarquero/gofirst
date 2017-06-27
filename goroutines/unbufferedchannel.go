package main

import (
  "fmt"
  "time"
)

func main()  {
  channel := make(chan string, 1)
  go func ()  {
    channel <- "Hello World! 1"
    channel <- "Hello World! 2"
    fmt.Println("Finishing goroutine")
  }()

  time.Sleep(time.Second)
  message := <- channel
  fmt.Println(message)

}
