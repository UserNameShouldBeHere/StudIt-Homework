package main

import (
	"time"
)

type User struct {
	Name       string
	Login      string
	Password   string
	LastSeenAt time.Time
}

type Admin struct {
	User
	Access uint8
}

// Здесь должна быть функция изменения имени пользователя
func renameUser(user interface{}, newName string) {
	switch user.(type) {
	case *User:
		user := user.(*User)
		user.Name = newName
	case *Admin:
		user := user.(*Admin)
		if user.Access < 3 {
			user.Name = newName
		}
	}
}

func main() {

}
