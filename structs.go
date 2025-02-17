package main

import (
	"fmt"

	"exmple.com/pointers/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		panic(err)
	}

	// ... do something awesome with that gathered data!
	// fmt.Print(appUser)

	appUser.ClearUserName()
	appUser.OutputUserDetailds()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
