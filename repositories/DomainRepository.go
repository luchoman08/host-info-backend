package repositories

import (
	"../app"
	"../interfaces"
	"../models"
	"github.com/luchoman08/ssllabs"
	"log"
	"net/url"
	"time"
)

// DomainRepository implements all the needed methods for access domain data
type DomainRepository struct {
	interfaces.SSLabsHandler
	interfaces.GoScraperHandler
	interfaces.GORMHandler
	interfaces.ServerService
}

// CreateDomain save in local storage a domain
func (repository *DomainRepository) CreateDomain(domain *models.DomainModel) {
	domain.SearchedAt = time.Now()
	domain.LastMajorChange = time.Now()
	repository.GetDB().Create(domain)
}

// GetDomainFromLocal receive a url and returns the matching domain stored locally if exists
func (repository *DomainRepository) GetDomainFromLocal(u url.URL) models.DomainModel {
	app.NormalizeURL(&u)
	return repository.GetByHostName(u.Hostname())
}

// ExistByHostName check if a domain exists locally based in its HostName
func (repository *DomainRepository) ExistByHostName(hostName string) bool {
	domain := models.DomainModel{}
	return !repository.GetDB().Where(models.DomainModel{HostName: hostName}).First(&domain).RecordNotFound()
}

// GetByHostName returns a domain related to the given hostName if exists locally, also append the
// locally sotred servers related to the domain found
func (repository *DomainRepository) GetByHostName(hostName string) (domain models.DomainModel) {
	domain = models.DomainModel{}
	repository.GetDB().Where(models.DomainModel{HostName: hostName}).First(&domain)
	domain.Servers = repository.GetServersOfDomain(&domain)
	return
}

func (repository *DomainRepository) populateServers(domain *models.DomainModel) {
	domain.Servers = repository.GetServersOfDomain(domain)

}

// UpdateSearchedTime update a domain SearchedAt in local storage to the current time
func (repository *DomainRepository) UpdateSearchedTime(domain *models.DomainModel) {
	domain.SearchedAt = time.Now()
	repository.GetDB().Save(domain)
}

// GetLastSearched returns the last searched domains stored locally ordered desc by its
// searched at property
func (repository *DomainRepository) GetLastSearched(limit int) (domains []models.DomainModel) {
	repository.GetDB().Limit(limit).Order("searched_at desc").Find(&domains)
	for index := range domains {
		repository.populateServers(&domains[index])
	}
	return
}

// UpdateDomain update all the fields of a domain in the local storage, only works
// if the domain given have the ID value in a existing value
func (repository *DomainRepository) UpdateDomain(domain *models.DomainModel) {
	repository.GetDB().Save(domain)
}

func (repository *DomainRepository) appendScrap(u url.URL, domain *models.DomainModel) error {
	var scrap, err = repository.Scrape(u, 5)
	if err != nil {
		return err

	}
	domain.Logo = app.NormalizePageIcoURL(scrap.Preview.Icon, u)
	domain.Title = scrap.Preview.Title
	return err
}

// GetDomainByHostNameUpdatedBefore returns a domain only if locally exists a domain updated before than the given date
func (repository *DomainRepository) GetDomainByHostNameUpdatedBefore(hostName string, t time.Time) (
	domain models.DomainModel,
	found bool) {
	found = false
	repository.GetDB().Where("last_major_change < ?", t).First(&domain)
	if domain.HostName != "" {
		repository.populateServers(&domain)
		found = true
	}
	return
}

// GetDomainFromExtern find for extern info about a domain than match with the given URL and return that.
// If the server does not exists locally this is stored locally.
// Also, if the servers returned does not exists locally, they are saved.
func (repository *DomainRepository) GetDomainFromExtern(u url.URL) (
	domain models.DomainModel,
	endPoints []ssllabs.Endpoint,
	err error) {
	app.NormalizeURL(&u)
	hostName := u.Hostname()
	domain.HostName = hostName
	var report, reportErr = repository.GetDetailedReport(hostName)
	if reportErr != nil {
		err = reportErr
		return
	}
	endPoints = report.Endpoints
	app.NormalizeURLWithScheme(&u, report.Protocol)
	scrapErr := repository.appendScrap(u, &domain) // this error is not fatal
	if scrapErr != nil {
		log.Printf("scrap error was ocurred with host %s", domain.HostName)
	}
	domain.IsDown = report.Status != repository.ReadyState()
	if repository.ExistByHostName(domain.HostName) {
		dbDomain := repository.GetDomainFromLocal(u)
		repository.UpdateSearchedTime(&dbDomain)
	}
	var sslGrades []string
	for _, endPoint := range report.Endpoints {
		sslGrades = append(sslGrades, endPoint.Grade)
	}
	domain.SslGrade = app.GetMinorSSLGradeFromList(sslGrades)
	return
}
