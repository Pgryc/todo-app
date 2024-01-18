package router

import (
	"todo-backend/cmd/api/resource/health"
	"todo-backend/cmd/api/resource/task"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/livez", health.Read)
	r.Route("/v1", func(r chi.Router) {
		taskAPI := &task.API{}
		r.Get("/tasks", taskAPI.List)
		r.Post("/tasks", taskAPI.Create)
		r.Get("/tasks/{id}", taskAPI.Read)
		r.Put("/tasks/{id}", taskAPI.Update)
	})
	return r
}
