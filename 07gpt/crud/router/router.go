package router

import (
	"crud/handlers"
	"crud/repo"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux { //chi.Mux is the main router type in chi package. It implements the http.Handler interface and is used to define routes and middleware for handling HTTP requests.
	// create a new chi router instance
	r := chi.NewRouter()

	// create an in-memory store and handler instance
	store := repo.NewStore()
	h := handlers.NewStudentHandler(store)

	// API routes
	r.Route("/api", func(r chi.Router) { // chi.Router is an interface that defines the methods for routing HTTP requests in the chi package.
		r.Get("/students", h.List)
		r.Get("/students/{id}", h.GetByID)
		r.Post("/students", h.Create)
		r.Put("/students/{id}", h.Update)
		r.Delete("/students/{id}", h.Delete)
	})

	return r
}
