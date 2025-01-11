package services

import (
	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/models"
)

type TaskInterface interface {
	CreateTask(taskDto dto.TaskDto) (*models.Task, error)
	GetTaskById(id uint) (*models.Task, error)
	ListTasks() ([]models.Task, error)
	UpdateTask(id uint, taskDto dto.TaskDto) (*models.Task, error)
	DeleteTask(id uint) error
}

type TaskService struct {}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) CreateTask(taskDto dto.TaskDto) (*models.Task, error) {

	task := models.Task{
		Name: taskDto.Name,
		Description: taskDto.Description,
		OwnerID: taskDto.OwnerID,
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) GetTaskById(id uint) (*models.Task, error) {
	var task models.Task
	if err := config.DB.Preload("Owner").First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := config.DB.Preload("Owner").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}


func (s *TaskService) UpdateTask(id uint, taskDto dto.TaskDto) (*models.Task, error) {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return nil, err
	}

	task.Name = taskDto.Name
	task.Description = taskDto.Description
	task.OwnerID = taskDto.OwnerID

	if err := config.DB.Save(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}


