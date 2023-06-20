package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)

	aNumber := 88

	var ptr = &aNumber
	fmt.Println("Value of pointer is:", ptr)
	fmt.Println("Value of actual pointer is:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of actual pointer is:", *ptr)
}
