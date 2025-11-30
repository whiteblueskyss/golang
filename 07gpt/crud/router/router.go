package router

import (
	"crud/handlers"
	"crud/repo"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// create an in-memory store and handler instance
	store := repo.NewStore()
	h := handlers.NewStudentHandler(store)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/students", h.List)
		r.Get("/students/{id}", h.GetByID)
		r.Post("/students", h.Create)
		r.Put("/students/{id}", h.Update)
		r.Delete("/students/{id}", h.Delete)
	})

	return r
}
