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

// model for cources -files
type Course struct {
	CourseId       string   `json:"courseid"`
	CourseName     string   `json:"coursename"`
	CoursePrice    int      `json:"price"`
	Author        *Author   `json:"author"`

}
type Author struct{
	Fullname  string     `json:"fullname`
	Website   string     `json:website`

}

var courses []Course


//middleware, helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId =="" && c.CourseName  ==""
	return c.CourseName == ""

	
}
func main() {
	fmt.Println("API - Golang")
	r := mux.NewRouter()
   //seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React", CoursePrice: 299, Author: &Author{Fullname: "John Doe", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 199, Author: &Author{Fullname: "Kiran", Website: "co.dev"}})
	courses = append(courses, Course{CourseId: "5", CourseName: "Angular", CoursePrice: 299, Author: &Author{Fullname: "Shyam Doe", Website: "go.dev"}})

// routing
r.HandleFunc("/", serveHome).Methods("GET")
r.HandleFunc("/courses", getAllCourses).Methods("GET")
r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
r.HandleFunc("/course", createOneCourse).Methods("POST")
r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")



// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
    fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab is from request
	params := mux.Vars(r)

	// loop throught courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	    
	}
	json.NewEncoder(w).Encode("No course found with given id") // if no match was found send this message
	return
	
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create a new course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//what about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	defer r.Body.Close()
   if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	 }
	//generate unique id, string
	// append couese into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
func updateOneCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")
	//first - grab id from request's paramteres
	params := mux.Vars(r)


	//loop, id remove , add with my ID
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

	//set the status code and send response for not found
	

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
   //loop, id remove , add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
}

}