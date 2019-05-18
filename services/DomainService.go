package services

import (
	"../app"
	"../interfaces"
	"../models"
	"net/url"
)
// DomainService implements the methods for access the domain repository
type DomainService struct {
	interfaces.DomainRepository
}
// GetDomain returns a domain info than match with a given route
func (service *DomainService) GetDomain(route string) (domain models.DomainModel, err error) {
	u, err := url.Parse(route)
	app.NormalizeURL(u)
	domain, err = service.GetDomainFromExtern(*u)
	return
}
// ServiceGetLastSearched returns the last searched domains and limit these by the limit
// argument, if the number of domains searched is less than the given limit, the existant
// domains searched are returned
func (service *DomainService) ServiceGetLastSearched(limit int) (domains []models.DomainModel) {
	domains = service.GetLastSearched(limit)
	return
}
