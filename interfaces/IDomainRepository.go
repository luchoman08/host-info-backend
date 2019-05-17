package interfaces

import (
	"../models"
	"net/url"
)

type IDomainRepository interface {
	GetDomainFromLocal(url url.URL) models.DomainModel
	GetDomainFromExtern(url url.URL) (models.DomainModel, error)
	ExistByHostName(string) bool
	CreateDomain(*models.DomainModel)
	GetByHostName(string) models.DomainModel
}