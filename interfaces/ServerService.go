package interfaces

import (
	"../models"
	"github.com/luchoman08/ssllabs"
)

type ServerService interface {
	ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool
	GetServer(models.DomainModel, ssllabs.Endpoint) (server models.ServerModel, err error)
	GetServersOfDomain(model *models.DomainModel) []models.ServerModel
}
