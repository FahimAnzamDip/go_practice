package main

import "fmt"

func main() {
	defer fmt.Println(adder(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))

	fmt.Println()

	fahim := User{"Fahim Anzam", 24, "fahimanzam9@gmailcom", false}
	fahim.IsVerified()
}

func adder(values ...int) int {
	var total int = 0

	for _, value := range values {
		total += value
	}

	return total
}

type User struct {
	Name     string
	Age      int
	email    string
	verified bool
}

func (user User) IsVerified() {
	if user.verified {
		fmt.Println("Verified")
	} else {
		fmt.Println("Unverified")
	}
}
