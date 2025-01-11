package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/services"
)

// TaskHandler is a struct that contains the methods to handle the Task model

type TaskHandler struct {
	taskService services.TaskInterface
}

func NewTaskHandler(taskService services.TaskInterface) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var taskDto dto.TaskDto

	if err := c.ShouldBindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	task, err := h.taskService.CreateTask(taskDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func (h *TaskHandler) GetTaskById(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("taskId"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}
	
	task, err := h.taskService.GetTaskById(uint(id))
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
		})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}

	var taskDto dto.TaskDto
	if err := c.ShouldBindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	task, err := h.taskService.UpdateTask(uint(id), taskDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.taskService.ListTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}
	response := h.taskService.DeleteTask(uint(id))
	if response != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
