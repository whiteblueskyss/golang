package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func chiServer() {
	r := chi.NewRouter() // create a new router instance. This router will handle incoming HTTP requests and route them to the appropriate handlers based on the defined routes.

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from chi!") // write "Hello from chi!" to the response writer
	})

	fmt.Println("Server running on 8080...")
	http.ListenAndServe(":8080", r)
}
