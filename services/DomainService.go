package services

import (
	"../interfaces"
	"../models"
	"net/url"
)

type DomainService struct {
	interfaces.IDomainRepository
}

func (service *DomainService) GetDomain(route string) (domain models.DomainModel, err error) {
	u, err := url.Parse(route)
	if service.ExistByHostName(u.Hostname()) {
		domain = service.GetDomainFromLocal(*u)
		return
	} else {
		domain, err = service.GetDomainFromExtern(*u)
		return
	}
}
