package main

import "fmt"

func main() {
	type User struct {
		Name     string
		Age      int
		email    string
		verified bool
	}

	fahim := User{"Fahim Anzam", 24, "fahimanzam9@gamil.com", false}

	fmt.Println(fahim)
	fmt.Printf("%+v\n\n", fahim)
	fmt.Printf("%v | %v", fahim.Name, fahim.email)
}
