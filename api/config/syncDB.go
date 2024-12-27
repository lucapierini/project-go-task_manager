package config

import "github.com/lucapierini/project-go-task_manager/models"

func SyncDB() {
	// DB.AutoMigrate(&models.Project{})
	// DB.AutoMigrate(&models.Task{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Role{})
}