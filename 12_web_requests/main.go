package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&payment_id=49fnfh2flsd"

func main() {
	// response, err := http.Get(url)
	// checkNilErr(err)

	// fmt.Printf("Response is of type: %T", response)

	// defer response.Body.Close()

	// databytes, err := io.ReadAll(response.Body)
	// checkNilErr(err)
	// content := string(databytes)

	// fmt.Println(content)

	result, err := url.Parse(myUrl)
	checkNilErr(err)

	qparams := result.Query()

	for _, value := range qparams {
		fmt.Println(value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=fahim",
	}

	newUrl := partsOfUrl.String()

	fmt.Println(newUrl)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
