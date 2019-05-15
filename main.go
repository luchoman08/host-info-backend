package main

import (
	"./api"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r := chi.NewRouter()
	r.Use(cors.Handler)

	r.Route("/api/v1/", func(r chi.Router) {
		r.Mount("/", api.Routes())
	})
	http.ListenAndServe(":3000", r)
}
