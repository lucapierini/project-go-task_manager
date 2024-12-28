package config

// import (
// 	"github.com/lucapierini/project-go-task_manager/config"
// 	"github.com/lucapierini/project-go-task_manager/models"
// )

// func CreateDefaultRoles(){
// 	// Create default roles
// 	role := models.Role{Name: "Administrador"}
// 	if err := config.DB.Create(&role).Error; err != nil {
// 		return nil, err
// 	}
// 	return &role, nil
// }


// func CreateRole(roleDto dto.RoleDto) (*models.Role, error){
// 	role := models.Role{Name: roleDto.Name}
// 	if err := config.DB.Create(&role).Error; err != nil {
// 		return nil, err
// 	}
// 	return &role, nil
// }