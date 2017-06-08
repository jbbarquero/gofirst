package main

import (
  "fmt"
)

// notifier is an interface that defined notification type behavior.
type notifier interface {
  notify()
}

type user struct {
  name string
  email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
  fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
  n.notify()
}

func main() {
  u := user{"Javier", "javier@correo.es"}

  sendNotification(&u)
}
