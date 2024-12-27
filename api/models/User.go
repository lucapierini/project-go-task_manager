package models

import (
	"gorm.io/gorm"
	"github.com/lucapierini/project-go-task_manager/responses"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Roles []Role `gorm:"many2many:user_roles"`
}