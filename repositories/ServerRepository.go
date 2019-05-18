package repositories

import (
	"../interfaces"
	"../models"
	"github.com/luchoman08/ssllabs"
)
// ServerRepository implements all methods for access the server info
type ServerRepository struct {
	interfaces.GORMHandler
	interfaces.WhoIsHandler
}
// CreateServer stores a given server
func (repository *ServerRepository) CreateServer(model *models.ServerModel) {
	repository.GetDB().Create(model)
}
// GetServersForDomain returns all the servers stored locally than are under the given domain
func (repository *ServerRepository) GetServersForDomain(domain *models.DomainModel) (servers []models.ServerModel) {
	repository.GetDB().Where(models.ServerModel{DomainID: domain.ID}).Find(&servers)
	return
}
// GetServerFromExtern find the server from external resources and stores it locally if not exists
func (repository *ServerRepository) GetServerFromExtern(endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	server = models.ServerModel{}
	server.SslGrade = endpoint.Grade
	if endpoint.ServerName != "" {
		server.ServerName = endpoint.ServerName
	} else {
		server.ServerName = endpoint.IPAddress
	}
	server.IPAddress = endpoint.IPAddress
	var whoIs, nonFatalErr = repository.GetWhoIsParsed(endpoint.IPAddress)
	err = nonFatalErr
	if err != nil {
		err = err
	} else {
		server.Country = whoIs["country"]
		// In some cases, orgname is empty, if is the case
		// check for descr value than tipically have the name
		// of the organization
		if whoIs["orgname"] != "" {
			server.Owner = whoIs["orgname"]
		} else if whoIs["descr"] != "" {
			server.Owner = whoIs["descr"]
		}
	}
	return server, err
}
// GetServerFromLocal find the server locally byy IPAddress and if it exists the it is returned
func (repository *ServerRepository) GetServerFromLocal(ipAddress string) (server models.ServerModel) {
	repository.GetDB().Where(models.ServerModel{IPAddress: ipAddress}, ipAddress).First(&server)
	return
}
// ExistsByIPAdress check if a exists a server stored locally than match with a given IPAdress
func (repository *ServerRepository) ExistsByIPAdress(ipAddress string) bool {
	server := &models.ServerModel{}
	return !repository.GetDB().Where(models.ServerModel{IPAddress: ipAddress}).First(&server).RecordNotFound()
}
