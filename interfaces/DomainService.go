package interfaces

import (
	"../models"
)

type DomainService interface {
	GetDomain(route string) (models.DomainModel, error)
	ServiceGetLastSearched(resultLimit int) []models.DomainModel
}