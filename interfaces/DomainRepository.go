package interfaces

import (
	"../models"
	"net/url"
)
// DomainRepository provide the methods needed for manage the
// data sources related to Domains
type DomainRepository interface {
	GetDomainFromLocal(url url.URL) models.DomainModel
	GetDomainFromExtern(url url.URL) (models.DomainModel, error)
	ExistByHostName(string) bool
	CreateDomain(*models.DomainModel)
	GetByHostName(string) models.DomainModel
	UpdateSearchedTime(domain *models.DomainModel)
	GetLastSearched(limit int) []models.DomainModel
}
