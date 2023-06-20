package main

import (
	"fmt"
	"io"
	"net/http"
)

const fakeUrl string = "https://dummyjson.com/products/1"

func main() {
	GetRequest()
}

func GetRequest() {
	response, err := http.Get(fakeUrl)
	checkNilErr(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	checkNilErr(err)
	jsonData := string(content)

	fmt.Println(jsonData)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
