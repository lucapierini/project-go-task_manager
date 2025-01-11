package services

import (
	"errors"

	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/models"
)

type ProjectInterface interface {
	CreateProject(projectDto dto.ProjectDto) (*models.Project, error)
	GetProjectById(id uint) (*models.Project, error)
	ListProjects() ([]models.Project, error)
	UpdateProject(id uint, projectDto dto.ProjectDto) (*models.Project, error)
	ListProjectsByUserId(userId uint) ([]models.Project, error)
	DeleteProject(id uint) error
	AddUserToProject(projectId uint, userId uint) error
	RemoveUserFromProject(projectId uint, userId uint) error
	AddTaskToProject(projectId uint, taskId uint) error
	RemoveTaskFromProject(projectId uint, taskId uint) error
}

type ProjectService struct{}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) CreateProject(projectDto dto.ProjectDto) (*models.Project, error) {
	var existingProject models.Project
	if result := config.DB.Where("name = ?", projectDto.Name).First(&existingProject); result.Error == nil {
		return nil, errors.New("project already exists")
	}

	project := models.Project{
		Name:    projectDto.Name,
		Budget:  projectDto.Budget,
		OwnerID: projectDto.OwnerID,
	}

	if len(projectDto.UsersIds) > 0 {
		var users []models.User
		if err := config.DB.Find(&users, projectDto.UsersIds).Error; err != nil {
			return nil, err
		}
		project.Users = users

	}

	if len(projectDto.TasksIds) > 0 {
		var tasks []models.Task
		if err := config.DB.Find(&tasks, projectDto.TasksIds).Error; err != nil {
			return nil, err
		}
		project.Tasks = tasks
	}

	if err := config.DB.Create(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectService) GetProjectById(id uint) (*models.Project, error) {
	var project models.Project
	if result := config.DB.Preload("Users").Preload("Tasks").Preload("Owner").First(&project, id); result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (s *ProjectService) ListProjects() ([]models.Project, error) {
	var projects []models.Project
	if result := config.DB.Preload("Users").Preload("Tasks").Preload("Owner").Find(&projects); result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (s *ProjectService) UpdateProject(id uint, projectDto dto.ProjectDto) (*models.Project, error) {
	project, err := s.GetProjectById(id)
	if err != nil {
		return nil, err
	}

	project.Name = projectDto.Name
	project.Budget = projectDto.Budget

	if len(projectDto.UsersIds) > 0 {
		var users []models.User
		if err := config.DB.Find(&users, projectDto.UsersIds).Error; err != nil {
			return nil, err
		}
		project.Users = users
	}

	if len(projectDto.TasksIds) > 0 {
		var tasks []models.Task
		if err := config.DB.Find(&tasks, projectDto.TasksIds).Error; err != nil {
			return nil, err
		}
		project.Tasks = tasks
	}

	if err := config.DB.Save(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) ListProjectsByUserId(userId uint) ([]models.Project, error) {
	var projects []models.Project
	if result := config.DB.Preload("Users").Preload("Tasks").Preload("Owner").Where("owner_id = ?", userId).Find(&projects); result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (s *ProjectService) DeleteProject(id uint) error {
	var project models.Project
	if result := config.DB.First(&project, id); result.Error != nil {
		return result.Error
	}

	if result := config.DB.Delete(&project); result.Error != nil {
		return result.Error
	}

	return nil
}


func (s *ProjectService) AddUserToProject(projectId uint,  userId uint) error {
	var project models.Project
	if result := config.DB.Preload("Users").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	for _, user := range project.Users {
		if user.ID == userId {
			return errors.New("user is already in project")
		}
	}

	var user models.User
	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	err := config.DB.Model(&project).Association("Users").Append(&user)

	if err != nil {
		return err
	}
	return nil
}

func (s *ProjectService) RemoveUserFromProject(projectId uint, userId uint) error {
	var project models.Project
	if result := config.DB.Preload("Users").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	// Check if the user is the owner or part of the project
	find := false
	for _, user := range project.Users {
		if user.ID == userId {
			find = true
		}
	}
	if !find {
		return errors.New("user is not in project")
	}

	var user models.User
	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	err := config.DB.Model(&project).Association("Users").Delete(&user)

	if err != nil {
		return err
	}
	return nil
}

func (s *ProjectService) AddTaskToProject(projectId uint, taskId uint) error {
	var project models.Project
	if result := config.DB.Preload("Tasks").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	for _, task := range project.Tasks {
		if task.ID == taskId {
			return errors.New("task is already in project")
		}
	}

	var task models.Task
	if err := config.DB.Where("id = ?", taskId).First(&task).Error; err != nil {
		return err
	}

	err := config.DB.Model(&project).Association("Tasks").Append(&task)

	if err != nil {
		return err
	}
	return nil
}

func (s *ProjectService) RemoveTaskFromProject(projectId uint, taskId uint) error {
	var project models.Project
	if result := config.DB.Preload("Tasks").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	find := false
	// Check if the task is already in the project
		for _, task := range project.Tasks {
			if task.ID == taskId {
				find = true
			}
		}

		if !find {
			return errors.New("task is not in project")
		}

		// Add the task to the project
		var task models.Task
		if err := config.DB.Where("id = ?", taskId).First(&task).Error; err != nil {
			return err
		}

		err := config.DB.Model(&project).Association("Tasks").Delete(&task)
		if err != nil {
			return err
		}

	return nil
}

