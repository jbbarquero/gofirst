package main

import (
  "time"
  "fmt"
)

func main()  {
  helloCh := make(chan string, 1)
  goodbyCh := make(chan string, 1)
  quitCh := make(chan bool)
  go receiver(helloCh, goodbyCh, quitCh)
  go sendString(helloCh, "World!")
  time.Sleep(time.Second)
  go sendString(goodbyCh, "American Pie")
  <-quitCh
}

func sendString(ch chan<- string, s string)  {
  ch <- s
}

func receiver(helloCh, goodbyCh <-chan string, quitCh chan<- bool)  {
  for {
    select {
    case msg := <-helloCh:
      fmt.Printf("Hello %s\n", msg)
    case msg := <-goodbyCh:
      fmt.Printf("Bye %s\n", msg)
    case <- time.After(time.Second * 2):
      println("Nothing received in 2 seconds. Exiting")
      quitCh <- true
      break
    }
  }

}
