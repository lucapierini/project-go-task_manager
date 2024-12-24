package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
	Active bool `gorm:"default:true"`
	CreatedAt string `gorm:"-"`
	UpdatedAt string `gorm:"-"`
	DeletedAt string `gorm:"-"`
}