package models

import (
	"time"
)

// DomainModel represents a web domain
type DomainModel struct {
	ID               uint          `json:"-"`
	CreatedAt        time.Time     `json:"-"`
	UpdatedAt        time.Time     `json:"-" sql:"DEFAULT:current_timestamp"`
	LastMajorChange  time.Time     `json:"-" sql:"DEFAULT:current_timestamp"`
	SearchedAt       time.Time     `json:"-"`
	HostName         string        `gorm:"unique;not null" json:"host_name"`
	DeletedAt        *time.Time    `sql:"index"  json:"-"`
	Servers          []ServerModel `json:"servers" gorm:"foreignkey:DomainID;association_autoupdate:false;association_autocreate:false;column:related_record"`
	ServersChanged   bool          `json:"servers_changed"`
	SslGrade         string        `json:"ssl_grade"`
	PreviousSslGrade string        `json:"previous_ssl_grade"`
	Logo             string        `json:"logo"`
	Title            string        `json:"title"`
	IsDown           bool          `json:"is_down"`
}
