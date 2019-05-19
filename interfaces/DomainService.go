package interfaces

import (
	"../models"
)
// DomainPagedResult represents a domain result when the query its paginated
type DomainPagedResult struct {
	Domains []models.DomainModel `json:"domains"`
	Pages int `json:"pages"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}

// DomainService provide an interface to acces the domain resources
// repositories
type DomainService interface {
	GetDomain(route string) (models.DomainModel, error)
	ServiceGetLastSearched(limit int, page int) DomainPagedResult
}
