package main

import (
	"log"
)

// Admin represents a type of user with privileges
type Admin struct {
	User
	Level string
}

// User represents a user with name and email
type User struct {
	Name  string
	Email string
}

// Notify logs user's name and email
func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

// Notify logs admin's name and email
func (a *Admin) Notify() error {
	log.Printf("Admin: Sending Admin Email To %s<%s>\n",
		a.Name,
		a.Email)
	return nil
}

// Notifier contains shared methods
type Notifier interface {
	Notify() error
}

// SendNotification receives a Notifier and logs a notification
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
