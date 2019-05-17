package models

import (
	"time"
)

type DomainModel struct {
	ID               uint          `json:"-"`
	CreatedAt        time.Time     `json:"-"`
	UpdatedAt        time.Time     `json:"-"`
	HostName         string        `gorm:"unique;not null" json:"-"`
	DeletedAt        *time.Time    `sql:"index"  json:"-"`
	Servers          []ServerModel `json:"servers"`
	ServersChanged   bool          `json:"servers_changed"`
	SslGrade         string        `json:"ssl_grade"`
	PreviousSslGrade string        `json:"previous_ssl_grade"`
	Logo             string        `json:"logo"`
	Title            string        `json:"title"`
	IsDown           bool          `json:"is_down"`
}
