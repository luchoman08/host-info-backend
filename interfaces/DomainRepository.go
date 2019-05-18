package interfaces

import (
	"../models"
	"net/url"
)

type DomainRepository interface {
	GetDomainFromLocal(url url.URL) models.DomainModel
	GetDomainFromExtern(url url.URL) (models.DomainModel, error)
	ExistByHostName(string) bool
	CreateDomain(*models.DomainModel)
	GetByHostName(string) models.DomainModel
	UpdateSearchedTime(domain *models.DomainModel)
	GetLastSearched(limit int) []models.DomainModel
}
