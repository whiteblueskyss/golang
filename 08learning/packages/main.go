package main

import (
	"fmt"
	"github.com/fatih/color"
	"learning/package/auth"
	"learning/package/user"
)

func main() {
	auth.LoginWithCredentials("golang", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}
