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

	fmt.Println("Enter number:")
	a, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)
	aConv, _ := strconv.ParseInt(a, 10, 32)

	fmt.Println("Enter number:")
	b, _ := reader.ReadString('\n')
	b = strings.TrimSpace(b)
	bConv, _ := strconv.ParseInt(b, 10, 32)

	fmt.Println("Enter number:")
	c, _ := reader.ReadString('\n')
	c = strings.TrimSpace(c)
	cConv, _ := strconv.ParseInt(c, 10, 32)

	if aConv > bConv && aConv > cConv {
		fmt.Println(aConv)
	} else if bConv > aConv && bConv > cConv {
		fmt.Println(bConv)
	} else {
		fmt.Println(cConv)
	}
}
