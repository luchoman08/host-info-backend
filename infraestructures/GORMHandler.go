package infraestructures

import "github.com/jinzhu/gorm"

// GORMHandler implements all methods to manage the gorm manager db
// and save an instance of this for share the instance across all the repositories
type GORMHandler struct {
	Db *gorm.DB
}

// GetDB returns the gorm db instance for share it across all the repositories
func (handler *GORMHandler) GetDB() *gorm.DB {
	return handler.Db
}
