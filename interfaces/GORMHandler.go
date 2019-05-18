package interfaces

import "github.com/jinzhu/gorm"

type GORMHandler interface {
	GetDB() *gorm.DB
}
