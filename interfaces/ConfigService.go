package interfaces

import (
	"../models"
)

type ConfigService interface {
	GetConfig() *models.Config
	ReadConfig() (err error)
}
