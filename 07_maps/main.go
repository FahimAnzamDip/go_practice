package main

import "fmt"

func main() {
	skills := make(map[int]string)

	skills[1] = "Javascript"
	skills[2] = "Vue Js"
	skills[3] = "Php"
	skills[4] = "Laravel"
	skills[5] = "Go"

	fmt.Println(skills)
	fmt.Println(skills[3])

	delete(skills, 5)

	fmt.Println(skills)

	for key, value := range skills {
		fmt.Printf("%v => %v\n", key, value)
	}
}
