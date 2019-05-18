package interfaces

import (
	"../models"
)
// ConfigService provide all the functions related to app config values,
// providing methods for access the config and init this.
type ConfigService interface {
	GetConfig() *models.Config
	ReadConfig() (err error)
}
