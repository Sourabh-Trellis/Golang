package main

import "fmt"

var actual_username string = "sourabh@trellissoft"
var actual_password string = "sdsd@1259"

func main() {

	var entered_username string
	var entered_password string

	fmt.Print("Enter your username:")
	fmt.Scan(&entered_username)

	fmt.Print("Enter your password:")
	fmt.Scan(&entered_password)

	if entered_username == actual_username {
		if entered_password == actual_password {
			fmt.Println("Login Succesfull")
		} else {
			fmt.Println("Invalid password")
		}
	} else {
		fmt.Println("Invalid username.")
	}
}
