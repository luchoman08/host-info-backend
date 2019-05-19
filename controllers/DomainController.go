package controllers

import (
	"../interfaces"
	"github.com/go-chi/render"
	"github.com/luchoman08/ssllabs"
	"net/http"
	"strconv"
)

// DomainController implements all methods for provide domain info to web server interfaces
type DomainController struct {
	interfaces.DomainService
}

// ControllerGetLastSearched return a web response with the last searched domains
func (controller DomainController) ControllerGetLastSearched(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := 3
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "", http.StatusUnprocessableEntity)
			return
		}
		limit = limitInt
	}
	domains := controller.ServiceGetLastSearched(limit)
	render.JSON(w, r, domains)
}

// ControllerGetServer return a web response with a domain info if is available
func (controller DomainController) ControllerGetServer(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Query().Get("host")

	if route == "" {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	domain, err := controller.GetDomain(route)

	if err != nil {
		switch err.(type) {
		case ssllabs.RetriesExeed:
			{
				http.Error(w, http.StatusText(http.StatusPartialContent), http.StatusPartialContent)
			}
		case ssllabs.UnableToResolveDomain:
			{
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}
		default:
			{
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
		}
		return
	}
	render.JSON(w, r, domain)
}
