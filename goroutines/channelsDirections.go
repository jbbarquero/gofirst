package main

import (
  "fmt"
  "time"
)

func main()  {
  channel := make (chan string, 1)
  go func (ch chan<- string)  {
    ch <- "Hello World!"
    fmt.Println("Finishing goroutine")
  }(channel)

  time.Sleep(time.Second)
  message := <- channel
  fmt.Println(message)

  channel <- "Goooood bye!"
  go receiving(channel)
  time.Sleep(time.Second)
}

func receiving(ch <-chan string)  {
  msg := <- ch
  fmt.Println(msg)
}
