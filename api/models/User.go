package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Active bool `gorm:"default:true"`
	CreatedAt string `gorm:"-"`
	UpdatedAt string `gorm:"-"`
	DeletedAt string `gorm:"-"`
}