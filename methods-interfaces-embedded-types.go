package main

import (
	//	"fmt"
	"log"
)

type Admin struct {
	User
	Level string
}

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func (a *Admin) Notify() error {
	log.Printf("Admin: Sending Admin Email To %s<%s>\n",
		a.Name,
		a.Email)
	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	admin := &Admin{
		User: User{
			Name:  "janet jones",
			Email: "janet@email.com",
		},
		Level: "super",
	}

	SendNotification(admin)
	admin.User.Notify()
	admin.Notify()
}
