package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	s := Student{
		ID:   1,
		Name: "Selim",
		Age:  25,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func server() {
	http.HandleFunc("/student", studentHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
