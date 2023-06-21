package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Course - file
type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("<========Fake API========>")

	//Declaring router
	router := mux.NewRouter()

	//Seeding the Fake DB
	seedDb()

	//Routing
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	router.HandleFunc("/course", storeCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//Start Server
	fmt.Println("API URL: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//Controller - file

// Handle Home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><========Fake API========><h1/>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "No course found with the given ID.", "ID": params["id"]})
}

func storeCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//some not empty checks
	if r.Body == nil {
		json.NewEncoder(w).Encode("The request is Empty!")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("The request is Empty!")
		return
	}

	//generate unique id, string
	//append course into courses

	seed := time.Now().UnixNano()
	randGenerator := rand.New(rand.NewSource(seed))
	course.CourseId = strconv.Itoa(randGenerator.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get the id
	params := mux.Vars(r)

	//loop, id, remove, add with same id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "No course found with the given ID.", "ID": params["id"]})
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get the id
	params := mux.Vars(r)

	//loop, id, remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Course deleted successfully!"})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "No course found with the given ID.", "ID": params["id"]})
}

// Fake DB Seeder
func seedDb() {
	// Append 1
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Introduction to Go Programming",
		CoursePrice: 49,
		Author: &Author{
			FullName: "John Smith",
			Website:  "https://www.example.com/johnsmith",
		},
	})

	// Append 2
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Advanced Data Structures in Go",
		CoursePrice: 69,
		Author: &Author{
			FullName: "Jane Doe",
			Website:  "https://www.example.com/janedoe",
		},
	})

	// Append 3
	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "Web Development with Go",
		CoursePrice: 59,
		Author: &Author{
			FullName: "Alex Johnson",
			Website:  "https://www.example.com/alexjohnson",
		},
	})

	// Append 4
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "Concurrency in Go",
		CoursePrice: 49,
		Author: &Author{
			FullName: "Emily Thompson",
			Website:  "https://www.example.com/emilythompson",
		},
	})

	// Append 5
	courses = append(courses, Course{
		CourseId:    "5",
		CourseName:  "Building RESTful APIs with Go",
		CoursePrice: 79,
		Author: &Author{
			FullName: "Michael Brown",
			Website:  "https://www.example.com/michaelbrown",
		},
	})
}
