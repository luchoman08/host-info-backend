package models

import (
	"time"
)

type ServerModel struct {
	ID          uint        `json:"-"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
	DeletedAt   *time.Time  `sql:"index"  json:"-"`
	ServerName	string      `gorm:"unique" json:"address"`
	IpAddress   string		`gorm:"unique; not null" json:"-"`
	SslGrade    string      `json:"ssl_grade"`
	Country     string      `json:"country"`
	Owner       string      `json:"owner"`
	DomainModel DomainModel `gorm:"foreignkey:DomainID"  json:"-"` // use DomainID as foreign key
	DomainID    uint        `json:"-"`
}
