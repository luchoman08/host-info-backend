package controllers

import (
	"../app"
	"../interfaces"
	"github.com/go-chi/render"
	"github.com/luchoman08/ssllabs"
	"net/http"
	"strconv"
)

type DomainController struct {
	interfaces.IDomainService
}

func (controller DomainController) ControllerGetLastSearched(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := 3
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		} else {
			limit = limitInt
		}
	}
 	domains := controller.ServiceGetLastSearched(limit)
	render.JSON(w, r, domains)
}

func (controller DomainController) ControllerGetServer(w http.ResponseWriter, r *http.Request) {
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
