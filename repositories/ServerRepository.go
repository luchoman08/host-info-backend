package repositories

import (
	"../interfaces"
	"../models"
	"github.com/luchoman08/ssllabs"
)

type ServerRepository struct {
	interfaces.IGORMHandler
	interfaces.WhoIsHandler
}

func (repository *ServerRepository) CreateServer(model *models.ServerModel) {
	repository.GetDB().Create(model)
}
func (repository *ServerRepository) GetServersForDomain(domain models.DomainModel) (servers []models.ServerModel) {
	repository.GetDB().Where(models.ServerModel{DomainID: domain.ID}).Find(&servers)
	return
}
func (repository *ServerRepository) GetServerFromExtern(endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	server = models.ServerModel{}
	server.SslGrade = endpoint.Grade
	if endpoint.ServerName != "" {
		server.ServerName = endpoint.ServerName
	} else {
		server.ServerName = endpoint.IPAddress
	}
	server.IpAddress = endpoint.IPAddress
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
func (repository *ServerRepository) GetServerFromLocal(ipAddress string) (server models.ServerModel) {
	repository.GetDB().Where(models.ServerModel{IpAddress: ipAddress}, ipAddress).First(&server)
	return
}
func (repository *ServerRepository) ExistsByIpAddress(ipAddress string) bool {
	server := &models.ServerModel{}
	return !repository.GetDB().Where(models.ServerModel{IpAddress: ipAddress}).First(&server).RecordNotFound()
}
