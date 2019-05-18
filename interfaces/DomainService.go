package interfaces

import (
	"../models"
)
// DomainService provide an interface to acces the domain resources
// repositories
type DomainService interface {
	GetDomain(route string) (models.DomainModel, error)
	ServiceGetLastSearched(resultLimit int) []models.DomainModel
}
