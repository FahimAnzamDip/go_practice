package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Instructor  string   `json:"instructor"`
	Description string   `json:"-"`
	Students    []string `json:"enrolled_student_names"`
}

func main() {
	courses := []Course{
		{
			ID:          1,
			Title:       "Introduction to Go",
			Instructor:  "John Doe",
			Description: "Learn the basics of Go programming language.",
			Students:    []string{"Alice", "Bob", "Charlie"},
		},
		{
			ID:          2,
			Title:       "Web Development",
			Instructor:  "Jane Smith",
			Description: "Build modern web applications using popular frameworks.",
			Students:    []string{"Eve", "Frank", "Grace"},
		},
		{
			ID:          3,
			Title:       "Data Science",
			Instructor:  "Alex Johnson",
			Description: "Explore data analysis and machine learning techniques.",
			Students:    []string{"Harry", "Isabella", "Jack"},
		},
		{
			ID:          4,
			Title:       "Mobile App Development",
			Instructor:  "Michael Brown",
			Description: "Create mobile apps for iOS and Android platforms.",
			Students:    []string{"Olivia", "Parker", "Quinn"},
		},
		{
			ID:          5,
			Title:       "Database Management",
			Instructor:  "Sophia Lee",
			Description: "Master the art of managing databases effectively.",
			Students:    []string{"Ryan", "Scarlett", "Taylor"},
		},
	}

	jsonDataBytes, err := json.MarshalIndent(courses, "", "\t")
	jsonData := string(jsonDataBytes)
	checkNilErr(err)

	fmt.Println(jsonData)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
