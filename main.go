package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/X1Zeth2X/projectify-api/api/project"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Routes initialize Chi Mux routes and settings
func Routes() *chi.Mux {
	r := chi.NewRouter()

	// Add CORS

	r.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-type headers to application/json
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		// Attach JWTAuth
	)

	// Mount projects route
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/api/project", project.Routes())
	})

	return r
}

func main() {
	r := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // Panic if there is an error
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Printf("Starting server on %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
