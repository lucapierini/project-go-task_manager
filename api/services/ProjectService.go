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
	AddUsersToProject(projectId uint, userIds []uint) error
	RemoveUsersFromProject(projectId uint, userIds []uint) error
	AddTasksToProject(projectId uint, taskIds []uint) error
	RemoveTasksFromProject(projectId uint, taskIds []uint) error
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
	if result := config.DB.Preload("Users").Preload("Tasks").First(&project, id); result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (s *ProjectService) ListProjects() ([]models.Project, error) {
	var projects []models.Project
	if result := config.DB.Preload("Users").Preload("Tasks").Find(&projects); result.Error != nil {
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
	if result := config.DB.Preload("Users").Preload("Tasks").Joins("Users").Where("user_id = ?", userId).Find(&projects); result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (s *ProjectService) DeleteProject(id uint, userId uint) error {
	var project models.Project
	if result := config.DB.First(&project, id); result.Error != nil {
		return result.Error
	}

	if project.OwnerID != userId {
		return errors.New("only the project owner can delete the project")
	}

	if result := config.DB.Delete(&project); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ProjectService) AddUsersToProject(projectId uint, userId uint, userIds []uint) error {
	var project models.Project
	if result := config.DB.Preload("Users").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	// Check if the user is the owner or part of the project
	isAuthorized := project.OwnerID == userId
	if !isAuthorized {
		for _, user := range project.Users {
			if user.ID == userId {
				isAuthorized = true
				break
			}
		}
	}

	if !isAuthorized {
		return errors.New("user is not authorized to add users to this project")
	}

	var users []models.User
	if result := config.DB.Find(&users, userIds); result.Error != nil {
		return result.Error
	}

	project.Users = append(project.Users, users...)

	if result := config.DB.Save(&project); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ProjectService) RemoveUsersFromProject(projectId uint, userId uint, userIds []uint) error {
	var project models.Project
	if result := config.DB.Preload("Users").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	// Check if the user is the owner or part of the project
	isAuthorized := project.OwnerID == userId
	if !isAuthorized {
		for _, user := range project.Users {
			if user.ID == userId {
				isAuthorized = true
				break
			}
		}
	}

	if !isAuthorized {
		return errors.New("user is not authorized to remove users from this project")
	}

	var users []models.User
	if result := config.DB.Find(&users, userIds); result.Error != nil {
		return result.Error
	}

	var newUsers []models.User
	for _, user := range project.Users {
		found := false
		for _, userId := range userIds {
			if user.ID == userId {
				found = true
				break
			}
		}
		if !found {
			newUsers = append(newUsers, user)
		}
	}

	project.Users = newUsers

	if result := config.DB.Save(&project); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ProjectService) AddTasksToProject(projectId uint, taskIds []uint, userId uint) error {

	var project models.Project
	if result := config.DB.Preload("Users").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	// Check if the user is the owner or part of the project
	isAuthorized := project.OwnerID == userId
	if !isAuthorized {
		for _, user := range project.Users {
			if user.ID == userId {
				isAuthorized = true
				break
			}
		}
	}

	if !isAuthorized {
		return errors.New("user is not authorized to remove users from this project")
	}

	var tasks []models.Task
	if result := config.DB.Find(&tasks, taskIds); result.Error != nil {
		return result.Error
	}

	project.Tasks = append(project.Tasks, tasks...)

	if result := config.DB.Save(&project); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ProjectService) RemoveTasksFromProject(projectId uint, taskIds []uint, userId uint) error {
	var project models.Project
	if result := config.DB.Preload("Users").First(&project, projectId); result.Error != nil {
		return result.Error
	}

	// Check if the user is the owner or part of the project
	isAuthorized := project.OwnerID == userId
	if !isAuthorized {
		for _, user := range project.Users {
			if user.ID == userId {
				isAuthorized = true
				break
			}
		}
	}

	if !isAuthorized {
		return errors.New("user is not authorized to remove users from this project")
	}

	var tasks []models.Task
	if result := config.DB.Find(&tasks, taskIds); result.Error != nil {
		return result.Error
	}

	var newTasks []models.Task
	for _, task := range project.Tasks {
		found := false
		for _, taskId := range taskIds {
			if task.ID == taskId {
				found = true
				break
			}
		}
		if !found {
			newTasks = append(newTasks, task)
		}
	}

	project.Tasks = newTasks

	if result := config.DB.Save(&project); result.Error != nil {
		return result.Error
	}

	return nil
}

