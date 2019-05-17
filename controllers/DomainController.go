package controllers

import (
	"../app"
	"../interfaces"
	"github.com/go-chi/render"
	"net/http"
)

type DomainController struct {
	interfaces.IDomainService
}

func (controller DomainController) GetServer(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Query().Get("host")

	if route == "" {
		http.Error(w, http.StatusText(400), 400)
		render.JSON(w, r, app.ErrHostCannotBeEmpty.Error())
		return
	}
	var domain, err = controller.GetDomain(route)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, domain)
}
