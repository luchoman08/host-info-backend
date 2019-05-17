package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	domainController := ServiceContainer().GetDomainController()

	r := chi.NewRouter()
	r.Use(cors.Handler)
	r.HandleFunc("/api/v1/analyze", domainController.ControllerGetServer)
	r.HandleFunc("/api/v1/lastSearched", domainController.ControllerGetLastSearched)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
