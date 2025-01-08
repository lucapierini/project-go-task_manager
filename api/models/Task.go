package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name string `gorm:"not null"`
	Description string
	Owner   User    `gorm:"foreignKey:OwnerID"`
	OwnerID uint
	Project Project `gorm:"foreignKey:ProjectID"`
	ProjectID uint
}