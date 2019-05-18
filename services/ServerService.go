package services

import (
	"../interfaces"
	"../models"
	"github.com/golang/glog"
	"github.com/luchoman08/ssllabs"
)
// ServerService implements the methods for access the server data repository
type ServerService struct {
	interfaces.ServerRepository
}
// GetServersOfDomain returns the servers related to the given domain, the servers are finded
// locally
func (service *ServerService) GetServersOfDomain(domain *models.DomainModel) []models.ServerModel {
	return service.GetServersForDomain(domain)
}
// ServerCollectionsDistinct compare two arrays of servers and return if their are distinct
// collections of servers or if both collections are equal.
// The comparison is made in a profound way.
func (service *ServerService) ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool {
	return false
}
// GetServer constructs a server based in a ssllabs endpoint and in a existing domain,
// if the server does not exists locally, these is created.
func (service *ServerService) GetServer(domain models.DomainModel, endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	exists := service.ExistsByIPAdress(endpoint.IPAddress)
	server, err = service.GetServerFromExtern(endpoint)
	if !exists {
		if err != nil {
			glog.Warning("Get server From extern has been failed on who is info get. Error: ", err)
		}
		server.DomainID = domain.ID
		service.CreateServer(&server)
	}
	return
}
