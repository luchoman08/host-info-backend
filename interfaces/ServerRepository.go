package interfaces

import (
	"../models"
	"github.com/luchoman08/ssllabs"
)

type ServerRepository interface {
	CreateServer(*models.ServerModel)
	GetServersForDomain(model models.DomainModel) (servers []models.ServerModel)
	GetServerFromExtern(endpoint ssllabs.Endpoint) (models.ServerModel, error)
	GetServerFromLocal(address string) (server models.ServerModel)
	ExistsByAddress(string) bool
}
