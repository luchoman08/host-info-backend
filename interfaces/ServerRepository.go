package interfaces

import (
	"../models"
	"github.com/luchoman08/ssllabs"
)

// ServerRepository provide the methods for manage the server info
type ServerRepository interface {
	CreateServer(*models.ServerModel)
	GetServersForDomain(model *models.DomainModel) (servers []models.ServerModel)
	GetServerFromExtern(endpoint ssllabs.Endpoint) (models.ServerModel, error)
	GetServerFromLocal(ipAddress string) (server models.ServerModel, found bool)
	ExistsByIPAdress(string) bool
	UpdateServer(server *models.ServerModel)
}
