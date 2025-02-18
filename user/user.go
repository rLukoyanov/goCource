package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (user *User) OutputUserDetailds() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "asd",
			lastName:  "asdasd",
			birthdate: "///",
			createdAt: time.Now(),
		},
	}
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}, nil
}
