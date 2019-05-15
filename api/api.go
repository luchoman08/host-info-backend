package api

import (
	"../app"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/luchoman08/ssllabs"
	"net/http"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/analyze", GetServerInfo)
	return router
}

func GetServerInfo(w http.ResponseWriter, r *http.Request) {
	var client, _ = ssllabs.NewClient()
	route := r.URL.Query().Get("host")
	if route == "" {
		http.Error(w, http.StatusText(400), 400)
		render.JSON(w, r, app.ErrHostCannotBeEmpty.Error())
		return
	}
	var domainInfo, err = app.ExtractDomainInfo(route, client)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, domainInfo)
}
