package services

import (
	"../interfaces"
	"../models"
	"../app"
		"net/url"
)

type DomainService struct {
	interfaces.IDomainRepository
}

func (service *DomainService) GetDomain(route string) (domain models.DomainModel, err error) {
	u, err := url.Parse(route)
	app.NormalizeUrl(u)
	domain, err = service.GetDomainFromExtern(*u)
	return
}
