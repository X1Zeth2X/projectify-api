package project

import (
	"github.com/X1Zeth2X/projectify-api/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

// Routes initialize project routes.
func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	})

	r.Get("/{projectID}", handlers.GetProject)
	r.Post("/create", handlers.CreateProject)
	r.Put("/{projectID}", handlers.UpdateProject)
	r.Delete("/{projectID}", handlers.DeleteProject)

	return r
}
