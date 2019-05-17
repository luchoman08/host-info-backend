package repositories

import (
	"../interfaces"
	"../models"
	"fmt"
	"github.com/luchoman08/ssllabs"
)

type ServerRepository struct {
	interfaces.IGORMHandler
	interfaces.WhoIsHandler
}

func (repository *ServerRepository) CreateServer(model *models.ServerModel) {
	fmt.Println(model.Address)
	repository.GetDB().Create(model)
}
func (repository *ServerRepository) GetServersForDomain(domain models.DomainModel) (servers []models.ServerModel) {
	repository.GetDB().Where(models.ServerModel{DomainID:domain.ID}).Find(&servers)
	return
}
func (repository *ServerRepository) GetServerFromExtern(endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	server = models.ServerModel{}
	server.SslGrade = endpoint.Grade
	server.Address = endpoint.ServerName
	var whoIs, nonFatalErr = repository.GetWhoIsParsed(endpoint.IPAddress)
	err = nonFatalErr
	if err != nil {
		err = err
	} else {
		server.Country = whoIs["Country"]
		server.Owner = whoIs["OrgName"]
	}
	return server, err
}
func (repository *ServerRepository) GetServerFromLocal(address string) (server models.ServerModel) {
	repository.GetDB().Where("address = ?", address).First(&server)
	return
}
func (repository *ServerRepository) ExistsByAddress(address string) bool {
	server := &models.ServerModel{}
	return !repository.GetDB().Where(models.ServerModel{Address: address}).First(&server).RecordNotFound()
}
