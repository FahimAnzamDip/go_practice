package main

import "fmt"

func main() {
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(weekdays); i++ {
		fmt.Println(weekdays[i])
	}

	fmt.Print("\n")

	for _, day := range weekdays {
		fmt.Println(day)
	}

	fmt.Print("\n")

	for _, day := range weekdays {
		if day == "Friday" {
			goto hurray
		} else {
			fmt.Println(day)
		}
	}

hurray:
	fmt.Println("IT'S FIRDAY!!!!!!")
}
