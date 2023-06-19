package main

import "fmt"

func main() {
	var username string = "John Doe"
	fmt.Println(username)
	fmt.Printf("Variable username is of type: %T \n\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable isLoggedIn is of type: %T \n\n", isLoggedIn)

	numberOfUsers := 45
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable numberOfUsers is of type: %T \n\n", numberOfUsers)
}
