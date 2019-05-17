package repositories

import (
	"../app"
	"../interfaces"
	"../models"
	"net/url"
)

type DomainRepository struct {
	interfaces.ISSLabsHandler
	interfaces.IGoScraperHandler
	interfaces.IGORMHandler
	interfaces.ServerService
}

func (repository *DomainRepository) CreateDomain(domain *models.DomainModel) {
	repository.GetDB().Create(domain)
}

func (repository *DomainRepository) GetDomainFromLocal(u url.URL) models.DomainModel {
	app.NormalizeUrl(&u)
	return repository.GetByHostName(u.Hostname())
}
func (repository *DomainRepository) ExistByHostName(hostName string) bool {
	domain := models.DomainModel{}
	return !repository.GetDB().Where(models.DomainModel{HostName: hostName}).First(&domain).RecordNotFound()
}
func (repository *DomainRepository) GetByHostName(hostName string) (domain models.DomainModel) {
	domain = models.DomainModel{}
	repository.GetDB().Where(models.DomainModel{HostName: hostName}).First(&domain)
	domain.Servers = repository.GetServersOfDomain(domain)
	return
}
func (repository *DomainRepository) GetDomainFromExtern(u url.URL) (domain models.DomainModel, err error) {
	app.NormalizeUrl(&u)
	hostName := u.Hostname()
	domain.HostName = hostName
	var report, report_err = repository.GetDetailedReport(hostName)
	if report_err != nil {
		err = report_err
		return
	}
	app.NormalizeUrlWithScheme(&u, report.Protocol)
	var scrap, scrap_err = repository.Scrape(u, 5)
	if scrap_err != nil {
		err = scrap_err
		return
	}
	domain.Logo = app.NormalizePageIcoUrl(scrap.Preview.Icon, u)
	domain.Title = scrap.Preview.Title
	if report.Endpoints != nil {
		domain.SslGrade = report.Endpoints[0].Grade
	}
	domain.IsDown = report.Status != repository.ReadyState()
	if !repository.ExistByHostName(domain.HostName) {
		repository.CreateDomain(&domain)
	}
	var servers []models.ServerModel
	for i := 0; i < len(report.Endpoints); i++ {
		server, _ := repository.GetServer(domain, report.Endpoints[i])
		servers = append(servers, server)
	}
	domain.Servers = servers
	return
}
