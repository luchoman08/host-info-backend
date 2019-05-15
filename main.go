package main

import (
	/*"fmt"
	"github.com/luchoman08/ssllabs"
	"./app"*/
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/luchoman08/ssllabs"
	"net/http"
	"./api"
)

type AppContext struct {
	SsllabsClli ssllabs.Client
}
func (users *AppContext) showHandler(w http.ResponseWriter, r *http.Request) {
	//now you can use users.db
}
func (users *AppContext) addHandler(w http.ResponseWriter, r *http.Request) {
	//now you can use users.db
}
func main() {

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	var client, _ = ssllabs.NewClient();
	appContext := &AppContext{SsllabsClli: *client}
	r := chi.NewRouter()
	r.Use(cors.Handler)

	r.Route("/api/v1/", func(r chi.Router) {
		r.Mount("/", api.Routes())
	})
	r.HandleFunc("/api/v1", appContext.showHandler)
	http.ListenAndServe(":3000", r)
	/*
	var domainInfo, _ = app.ExtractDomainInfo("https://google.com/", c)
	fmt.Println(domainInfo)*/
}
