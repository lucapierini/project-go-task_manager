package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name    string  `gorm:"unique;not null"`
	Budget  float64 `gorm:"not null"`
	Owner   User    `gorm:"foreignKey:OwnerID"`
	OwnerID uint
	Users   []User `gorm:"many2many:project_users"`
	Tasks   []Task `gorm:"many2many:project_tasks"`
}

// AddUsersToProject(projectId uint, userIds []uint) error
// RemoveUsersFromProject(projectId uint, userIds []uint) error
// AddTasksToProject(projectId uint, taskIds []uint) error
// RemoveTasksFromProject(projectId uint, taskIds []uint) error
