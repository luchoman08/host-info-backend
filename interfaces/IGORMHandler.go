package interfaces

import "github.com/jinzhu/gorm"

type IGORMHandler interface {
	GetDB() *gorm.DB
}
