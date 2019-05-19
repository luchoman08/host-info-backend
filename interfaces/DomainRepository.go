package interfaces

import (
	"../models"
	"github.com/luchoman08/ssllabs"
	"net/url"
	"time"
)

// DomainRepository provide the methods needed for manage the
// data sources related to Domains
type DomainRepository interface {
	GetDomainFromLocal(url url.URL) models.DomainModel
	GetDomainFromExtern(url url.URL) (models.DomainModel, []ssllabs.Endpoint, error)
	ExistByHostName(string) bool
	CreateDomain(*models.DomainModel)
	UpdateDomain(*models.DomainModel)
	GetDomainByHostNameUpdatedBefore(hostName string, t time.Time) (domain models.DomainModel, found bool)
	GetByHostName(string) models.DomainModel
	UpdateSearchedTime(domain *models.DomainModel)
	GetLastSearched(limit int) []models.DomainModel
}
