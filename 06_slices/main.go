package main

import "fmt"

func main() {
	var fruitList = []string{}

	fruitList = append(fruitList, "Apple", "Banana", "Watermelon", "Mango")
	fmt.Println(fruitList)

	var index int = 2
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)
}
