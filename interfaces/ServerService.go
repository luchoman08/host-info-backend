package interfaces

import (
	"../models"
	"github.com/luchoman08/ssllabs"
)

// ServerService provide the methods for access the server repository
type ServerService interface {
	ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool
	GetMinorSSLGrade([]models.ServerModel) string
	GenerateServer(models.DomainModel, ssllabs.Endpoint) (server models.ServerModel, err error)
	GetServersOfDomain(model *models.DomainModel) []models.ServerModel
	EqualServers(s1, s2 models.ServerModel) bool
	EqualSetOfServers(set1, set2 []models.ServerModel) bool
	ServiceCreateServer(model *models.ServerModel)
	GetServerByIP(ip  string) models.ServerModel
	UpdateLocalServersIfChanged([]models.ServerModel)
}
