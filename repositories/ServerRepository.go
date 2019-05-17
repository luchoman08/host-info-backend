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
func (repository *ServerRepository) GetServersForDomain(model models.DomainModel) (servers []models.ServerModel) {
	repository.GetDB().Where("domain_id = ?", model.ID).Find(&servers)
	return
}
func (repository *ServerRepository) GetServerFromExtern(endpoint ssllabs.Endpoint) (models.ServerModel , error) {
	server := models.ServerModel{}
	server.SslGrade = endpoint.Grade
	server.Address = endpoint.ServerName
	var whoIs, err = repository.GetWhoIsParsed(endpoint.IPAddress)
	if err != nil {
		return server, err
	}
	server.Country = whoIs["Country"]
	server.Owner = whoIs["OrgName"]
	return server, err
}
func (repository *ServerRepository) GetServerFromLocal(address string) (server models.ServerModel) {
	repository.GetDB().Where("address = ?", address).First(&server)
	return
}
func (repository *ServerRepository) ExistsByAddress(address string) bool {
	const undefined = "undefined"
	server := &models.ServerModel{}
	repository.GetDB().Where("address = ?", address).FirstOrInit(&server, &models.ServerModel{Address: undefined})
	return server.Address != undefined
}