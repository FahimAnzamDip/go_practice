package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first number:")
	a, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)
	convA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter second number:")
	b, _ := reader.ReadString('\n')
	b = strings.TrimSpace(b)
	convB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result float64 = convA + convB
	fmt.Println("\nResult is:", result)
}
