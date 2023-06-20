package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/about", aboutHandler)

	fmt.Println("Server listening on port 8080")
	fmt.Println("visit: http://localhost:8080/")
	http.ListenAndServe(":8080", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Page</h1> <a href=\"/about\">About Us</a>"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>About Page</h1> <a href=\"/\">Home</a>"))
}
