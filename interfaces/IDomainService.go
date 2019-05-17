package interfaces

import (
	"../models"
)

type IDomainService interface {
	GetDomain(route string) (models.DomainModel, error)
	ServiceGetLastSearched(resultLimit int) []models.DomainModel
}
