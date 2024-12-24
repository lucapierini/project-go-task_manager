package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status StatusEnum
	ProjectID uint
	Project Project `gorm:"foreignKey:ProjectID"`
}

type StatusEnum string

const (
	StatusPending   StatusEnum = "Pending"
	StatusCompleted StatusEnum = "Completed"
	StatusInProgress StatusEnum = "In Progress"
)

