package interfaces

import (
	"../models"
	"github.com/luchoman08/ssllabs"
)
// ServerService provide the methods for access the server repository
type ServerService interface {
	ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool
	GetMinorSSLGrade([]models.ServerModel) string
	GetServer(models.DomainModel, ssllabs.Endpoint) (server models.ServerModel, err error)
	GetServersOfDomain(model *models.DomainModel) []models.ServerModel
}
