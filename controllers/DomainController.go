package controllers

import (
	"../app"
	"../interfaces"
	"github.com/go-chi/render"
	"github.com/luchoman08/ssllabs"
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
	domain, err := controller.GetDomain(route)

	if err != nil {
		switch err.(type) {
		case ssllabs.RetriesExeed:
			{
				http.Error(w, http.StatusText(http.StatusPartialContent), http.StatusPartialContent)
				render.JSON(w, r, err.Error())
			}
		default:
			{
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				render.JSON(w, r, err.Error())
			}
		}
		return
	} else {
		render.JSON(w, r, domain)
	}
}