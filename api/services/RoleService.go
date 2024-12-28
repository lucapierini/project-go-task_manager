package services

import (
	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/models"
)

type RoleInterface interface {
	CreateRole(roleDto dto.RoleDto) (*models.Role, error)
	GetRoleById(id uint) (*models.Role, error)
	UpdateRole(id uint, roleDto dto.RoleDto) (*models.Role, error)
	DeleteRole(id uint) error
	ListRoles() ([]models.Role, error)
	}

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) CreateRole(roleDto dto.RoleDto) (*models.Role, error){
	role := models.Role{Name: roleDto.Name}
	if err := config.DB.Create(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *RoleService) GetRoleById(id uint) (*models.Role, error){
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *RoleService) UpdateRole(id uint, roleDto dto.RoleDto) (*models.Role, error){
	role, err := s.GetRoleById(id)
	if err != nil {
		return nil, err
	}
	role.Name = roleDto.Name
	if err := config.DB.Save(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) DeleteRole(id uint) error {
	if err := config.DB.Delete(&models.Role{},id).Error; err != nil {
		return err
	}
	return nil
}

func (s *RoleService) ListRoles()([]models.Role, error){
	var roles []models.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}