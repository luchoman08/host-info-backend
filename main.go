package main

import (
	/*"fmt"
	"github.com/luchoman08/ssllabs"
	"./app"*/
	"github.com/go-chi/chi"
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
	var client, _ = ssllabs.NewClient();
	appContext := &AppContext{SsllabsClli: *client}
	r := chi.NewRouter()
	r.Route("/api/v1/", func(r chi.Router) {
		r.Mount("/", api.Routes())
	})
	r.HandleFunc("/api/v1", appContext.showHandler)
	http.ListenAndServe(":3000", r)
	/*
	var domainInfo, _ = app.ExtractDomainInfo("https://google.com/", c)
	fmt.Println(domainInfo)*/
}
