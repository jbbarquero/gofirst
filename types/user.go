package main

import (
	"fmt"
)

// notifier is an interface that defined notification type behavior.
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

// notify implements a method that can be called via a value of type admin
func (a *admin) notify() {
	fmt.Printf("Sending ADMIN email to %s<%s>, level: %s\n",
		a.name, a.email, a.level)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Javier", "javier@correo.es"}

	sendNotification(&u)

	ad := admin{
		user: user{
			name:  "Admin",
			email: "admin@correo.es",
		},
		level: "super-power",
	}

	ad.user.notify()

	sendNotification(&ad)

}
