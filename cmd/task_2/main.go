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
	switch typedUser := user.(type) {
	case *User:
		typedUser.Name = newName
	case *Admin:
		if typedUser.Access < 3 {
			typedUser.Name = newName
		}
	}
}

func main() {

}
