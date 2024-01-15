package router

import (
	"todo-backend/cmd/api/resource/health"
	"todo-backend/cmd/api/resource/task"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/livez", health.Read)
	r.Route("/v1", func(r chi.Router) {
		itemAPI := &task.API{}
		r.Get("/items", itemAPI.List)
		r.Post("/items", itemAPI.Create)
		r.Get("/items/{id}", itemAPI.Read)
		r.Put("/items/{id}", itemAPI.Update)
	})
	return r
}
