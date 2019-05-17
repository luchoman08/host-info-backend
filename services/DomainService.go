package services

import (
	"../app"
	"../interfaces"
	"../models"
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

func (service *DomainService) ServiceGetLastSearched(limit int) (domains []models.DomainModel) {
	domains = service.GetLastSearched(limit)
	return
}