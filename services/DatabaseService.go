package services

import (
	"../interfaces"
)

// DatabaseService implements the methods for manage the database
type DatabaseService struct {
	interfaces.GORMHandler
	Interfaces []interface{}
}

// Migrate creates the initial tables defined by the Interfaces
func (service *DatabaseService) Migrate() {
	for _, element := range service.Interfaces {
		service.GetDB().AutoMigrate(element)
	}
}

// DropTables remove all databases previously created if this exists
func (service *DatabaseService) DropTables() {
	for _, element := range service.Interfaces {
		service.GetDB().DropTable(element)
	}
}
