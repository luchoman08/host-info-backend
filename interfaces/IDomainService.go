package interfaces

import (
	"../models"
)

type IDomainService interface {
	GetDomain(route string) (models.DomainModel, error)
}
