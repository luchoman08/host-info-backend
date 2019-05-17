package services

import (
	"../interfaces"
	"../models"
	"github.com/golang/glog"
	"github.com/luchoman08/ssllabs"
)

type ServerService struct {
	interfaces.ServerRepository
}

func (service *ServerService) GetServersOfDomain(domain *models.DomainModel) []models.ServerModel {
	return service.GetServersForDomain(domain)
}

func (service *ServerService) ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool {
	return false
}
func (service *ServerService) GetServer(domain models.DomainModel, endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	exists := service.ExistsByIpAddress(endpoint.IPAddress)
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
