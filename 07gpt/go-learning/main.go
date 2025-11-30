package main

// // package main tells Go:
// // “This file is part of a program that can be run as an executable.”
// // Only a package named main can produce a runnable program (a binary).
// // Packages with other names (like package utils, package models) are libraries.

// import (
// 	"encoding/json"
// 	"net/http"
// ) // for formating and printing text

func main() { // go alaways starts execution from main function.
	// errorDemo(10, 0)
	// http.HandleFunc("/students", studentsHandler)

	// server()

	chiServer()
}

// func studentsHandler(w http.ResponseWriter, r *http.Request) {
// 	students := []Student{
// 		{ID: 1, Name: "Selim", Age: 25},
// 		{ID: 2, Name: "Aziz", Age: 22},
// 		{ID: 3, Name: "Karim", Age: 20},
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(students)
// }

// func createStudentHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var s Student
// 	err := json.NewDecoder(r.Body).Decode(&s)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(s)
// }

// mux := http.NewServeMux()

// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "Hello!")
// })

// mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "You're inside /api route")
// })

// http.ListenAndServe(":8080", mux)
