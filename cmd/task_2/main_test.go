package main

import (
	"reflect"
	"testing"
)

func TestRenameUser(t *testing.T) {
	userCases := []struct {
		testName string
		user     User
		newName  string
		expected User
	}{
		{
			"change user name",
			User{},
			"test",
			User{
				Name: "test",
			},
		},
		{
			"change user name",
			User{
				Name:     "Ahmed",
				Login:    "ahmed@gmail.com",
				Password: "ratata",
			},
			"Nikita",
			User{
				Name:     "Nikita",
				Login:    "ahmed@gmail.com",
				Password: "ratata",
			},
		},
	}
	adminCases := []struct {
		testName string
		user     Admin
		newName  string
		expected Admin
	}{
		{
			"change admin name with access < 3",
			Admin{
				User: User{
					Name:     "Admin",
					Login:    "admin@mail.ru",
					Password: "root",
				},
				Access: 2,
			},
			"Pamparam",
			Admin{
				User: User{
					Name:     "Pamparam",
					Login:    "admin@mail.ru",
					Password: "root",
				},
				Access: 2,
			},
		},
		{
			"change admin name with access > 3",
			Admin{
				User: User{
					Name:     "Admin",
					Login:    "admin@mail.ru",
					Password: "root",
				},
				Access: 5,
			},
			"Pamparam",
			Admin{
				User: User{
					Name:     "Admin",
					Login:    "admin@mail.ru",
					Password: "root",
				},
				Access: 5,
			},
		},
	}

	for _, currentCase := range userCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			renameUser(&currentCase.user, currentCase.newName)

			if !reflect.DeepEqual(currentCase.user, currentCase.expected) {
				t.Errorf("Incorrect result: expected %v, got %v", currentCase.expected, currentCase.user)
			}
		})
	}

	for _, currentCase := range adminCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			renameUser(&currentCase.user, currentCase.newName)

			if !reflect.DeepEqual(currentCase.user, currentCase.expected) {
				t.Errorf("Incorrect result: expected %v, got %v", currentCase.expected, currentCase.user)
			}
		})
	}
}
