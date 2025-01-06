package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name string `gorm:"not null"`
	Description string
	Project Project `gorm:"foreignKey:ProjectID"`
	ProjectID uint
}