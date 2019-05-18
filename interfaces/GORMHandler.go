package interfaces

import "github.com/jinzhu/gorm"
// GORMHandler provide the methods for get the gorm instance
// across the entire application
type GORMHandler interface {
	GetDB() *gorm.DB
}
