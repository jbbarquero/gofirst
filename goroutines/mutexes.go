package main

import (
  "sync"
  "time"
)

type Counter struct {
  sync.Mutex
  value int
}

/*
* go run -race to see potential race conditions
*/

func main()  {
  counter := Counter{}
  for i := 0; i < 10; i++ {
    go func (i int)  {
      counter.Lock() //comment to see race conditions
      counter.value++
      defer counter.Unlock() //comment to see race conditions
    }(i)
  }
  time.Sleep(time.Second)
  counter.Lock()
  defer counter.Unlock()
  println(counter.value)

}
