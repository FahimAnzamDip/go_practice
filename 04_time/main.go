package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	formattedTime := currentTime.Format("01-02-2006 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)
}
