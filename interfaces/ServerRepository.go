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
	GetServerFromLocal(address string) (server models.ServerModel)
	ExistsByIPAdress(string) bool
}
