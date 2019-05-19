package services

import (
	"../app"
	"../interfaces"
	"../models"
	"net/url"
	"time"
)

// DomainService implements the methods for access the domain repository
type DomainService struct {
	interfaces.DomainRepository
	interfaces.ServerService
}

func (service *DomainService) updateOrSaveLocalDomainIfChanged(domain *models.DomainModel) (updated bool, prevGrade string) {
	updated = false
	prevGrade = ""
	oneHourBefore := time.Now().Add(-1 * time.Hour)
	existsLocally := service.ExistByHostName(domain.HostName)
	localDomainOneHourAgo, foundFromOneOurAgo := service.GetDomainByHostNameUpdatedBefore(domain.HostName, oneHourBefore)
	if !existsLocally {
		service.CreateDomain(domain)
		for _, server := range domain.Servers {
			server.DomainID = domain.ID
			service.ServiceCreateServer(&server)
		}
	}
	if foundFromOneOurAgo {
		serversChanged := !service.EqualSetOfServers(domain.Servers, localDomainOneHourAgo.Servers)
		if serversChanged {
			domain.ID = localDomainOneHourAgo.ID
			domain.LastMajorChange = time.Now()
			service.UpdateDomain(domain)
			prevServers := service.GetServersOfDomain(domain)
			service.UpdateLocalServersIfChanged(domain.Servers)
			prevGrade = service.GetMinorSSLGrade(prevServers)
			updated = true
		}
	}
	return
}

// GetDomain returns a domain info than match with a given route
// If the domain retrieved from external resources have a local copy
// and it has been changed, the local data of domain and its servers  are updated
func (service *DomainService) GetDomain(route string) (models.DomainModel, error) {
	u, urlErr := url.Parse(route)
	if urlErr != nil {
		return models.DomainModel{}, urlErr
	}
	app.NormalizeURL(u)
	domain, endPoints, err := service.GetDomainFromExtern(*u)
	if err != nil {
		return models.DomainModel{}, err
	}
	var servers []models.ServerModel
	for _, endPoint := range endPoints {
		server, _ := service.GenerateServer(domain, endPoint)
		servers = append(servers, server)
	}
	domain.Servers = servers
	updated, prevGrade := service.updateOrSaveLocalDomainIfChanged(&domain)
	if updated {
		domain.ServersChanged = true
		domain.PreviousSslGrade = prevGrade
	}
	return domain, err
}

// ServiceGetLastSearched returns the last searched domains and limit these by the limit
// argument, if the number of domains searched is less than the given limit, the existant
// domains searched are returned
func (service *DomainService) ServiceGetLastSearched(limit int) (domains []models.DomainModel) {
	domains = service.GetLastSearched(limit)
	return
}
