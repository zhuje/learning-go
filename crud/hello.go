package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Uppercase Field Names: Makes fields exported and accessible to external packages, including json encoding and decoding.
type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var students = []Person{
	{
		Id:   1,
		Name: "Abigail",
	},
	{
		Id:   2,
		Name: "Bob",
	},
	{
		Id:   3,
		Name: "Charlie",
	},
}

func findStudent(id int) (*Person, int) {
	for i := range students {
		if students[i].Id == id {
			return &students[i], i
		}
	}

	return nil, -1
}

// curl -X POST -H "Content-Type: application/json" -d "{\"id\": 4, \"name\": \"Ha Minh Nam\"}" localhost:8080/api/students
func createStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request")

	var person Person

	// get request body
	decoder := json.NewDecoder(r.Body)

	// Decode
	err := decoder.Decode(&person)

	fmt.Println(person)

	// handle error
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		fmt.Println("PUT request error while decoding JSON ")
		return
	}

	// handle success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"status": "success", "name": person.Name, "id": fmt.Sprintf("%d", person.Id)})

}

// curl -X GET localhost:8080/api/students | jq
func readStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET request\n", students)

	// handle success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if encodingError := json.NewEncoder(w).Encode(students); encodingError != nil {
		http.Error(w, fmt.Sprintf("error encoding JSOP %v", encodingError), http.StatusInternalServerError)
	}
}

// curl -X PATCH -H "Content-Type: application/json" -d "{\"id\": 3, \"name\": \"Dan\"}" localhost:8080/api/students
func updateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var updateInfo Person

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("PATCH request \n", updateInfo)

	student, _ := findStudent(updateInfo.Id)
	if student == nil {
		http.Error(w, fmt.Sprintf("Student with ID %d not found", updateInfo.Id), http.StatusNotFound)
	}

	if updateInfo.Name != "" {
		student.Name = updateInfo.Name
	}

	// handle success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if encodingError := json.NewEncoder(w).Encode(students); encodingError != nil {
		http.Error(w, fmt.Sprintf("error encoding JSOP %v", encodingError), http.StatusInternalServerError)
	}
}

// curl -X DELETE -H "Content-Type: application/json" -d "{\"id\": 3, \"name\": \"Charlie\"}" localhost:8080/api/students
func deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	var deleteInfo Person

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deleteInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("DELETE request \n", deleteInfo)

	// find index of the student
	_, index := findStudent(deleteInfo.Id)
	if index == -1 {
		http.Error(w, fmt.Sprintf("Student with ID %d not found", deleteInfo.Id), http.StatusNotFound)
		return
	}

	// cut the index out of the array
	students = append(students[:index], students[index+1:]...)

	if encodingError := json.NewEncoder(w).Encode(students); encodingError != nil {
		http.Error(w, fmt.Sprintf("error encoding JSOP %v", encodingError), http.StatusInternalServerError)
	}

}

func main() {

	port := ":8080"
	r := mux.NewRouter()
	fmt.Printf("Serving at port %s\n", port)

	r.HandleFunc("/api/students", createStudentHandler).Methods("POST")   // CREATE
	r.HandleFunc("/api/students", readStudentHandler).Methods("GET")      // READ
	r.HandleFunc("/api/students", updateStudentHandler).Methods("PATCH")  // UPDATE
	r.HandleFunc("/api/students", deleteStudentHandler).Methods("DELETE") // DELETE

	http.ListenAndServe(port, r)
}
