package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Budget  uint   `gorm:"not null"`
	Owner   User   `gorm:"foreignKey:OwnerID"`
	OwnerID uint
	Users   []User `gorm:"many2many:project_users"`
	Tasks   []Task `gorm:"many2many:project_tasks"`
}
