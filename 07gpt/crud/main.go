package main

import (
	"crud/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
