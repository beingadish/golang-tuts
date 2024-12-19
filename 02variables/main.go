package main

import "fmt"

const LoginToken string = "ftghbnjtfc" // Public in nature

func main() {
	var myName string = "Aadarsh"
	fmt.Println(myName)
	fmt.Printf("Variable is of type: %T \n", myName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// Another way of declaring variables
	loginCreds := 200
	fmt.Printf("Variable %d is of type: %T \n", loginCreds, loginCreds)

	// Accessing public variable
	fmt.Println(LoginToken)
}
