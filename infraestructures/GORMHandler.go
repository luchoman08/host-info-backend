package infraestructures

import "github.com/jinzhu/gorm"

type GORMHandler struct {
	Db *gorm.DB
}

func (handler *GORMHandler) GetDB() *gorm.DB {
	return handler.Db
}
