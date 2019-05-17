package services

import (
	"../interfaces"
)

type DatabaseService struct {
	interfaces.IGORMHandler
	Interfaces []interface{}
}

func (service *DatabaseService) Migrate()  {
	for _, element := range service.Interfaces {
		service.GetDB().AutoMigrate(element)
	}
}

func (service *DatabaseService) DropTables() {
	for _, element := range service.Interfaces {
		service.GetDB().DropTable(element)
	}
}