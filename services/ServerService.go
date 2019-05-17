package services

import (
	"../interfaces"
	"../models"
	"fmt"
	"github.com/luchoman08/ssllabs"
)

type ServerService struct {
	interfaces.ServerRepository
}

func (service *ServerService) ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool {
	return false
}
func (service *ServerService) GetServer(domain models.DomainModel, endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	exists := service.ExistsByAddress(endpoint.ServerName)
	fmt.Println(exists)
	if !exists {
		server, err = service.GetServerFromExtern(endpoint)
		if err != nil {
			return
		}
		server.DomainID = domain.ID
		service.CreateServer(&server)
		return
	} else {
		server = service.GetServerFromLocal(endpoint.ServerName)
		return
	}
}