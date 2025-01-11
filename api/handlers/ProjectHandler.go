package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/services"
)

type ProjectHandler struct {
	projectService services.ProjectInterface
}

func NewProjectHandler(projectService services.ProjectInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService : projectService,
	}
}

func (h *ProjectHandler) CreateProject(c *gin.Context){
	var projectDto dto.ProjectDto

	if err := c.ShouldBindJSON(&projectDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	project, err := h.projectService.CreateProject(projectDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"project": project,
	})
}

func (h *ProjectHandler) GetProjectById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("projectId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	project, err := h.projectService.GetProjectById(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (h *ProjectHandler) ListProjects(c *gin.Context){
	projects, err := h.projectService.ListProjects()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context){
	id, err := strconv.Atoi(c.Param("projectId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	var projectDto dto.ProjectDto
	if err := c.ShouldBindJSON(&projectDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	project, err := h.projectService.UpdateProject(uint(id), projectDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (h *ProjectHandler) ListProjectsByUserId(c *gin.Context){
	id, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
	}

	projects, err := h.projectService.ListProjectsByUserId(uint(id))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"projects": projects,
		"id_user": id,
	})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context){
	id, err := strconv.Atoi(c.Param("projectId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	err = h.projectService.DeleteProject(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project deleted successfully",
	})

}

func (h *ProjectHandler) AddUserToProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	idUser, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	err = h.projectService.AddUserToProject(uint(idProject), uint(idUser))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
		}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Users added to project successfully",
	})
}

func (h *ProjectHandler) RemoveUserFromProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
		return
	}

	idUser, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	err = h.projectService.RemoveUserFromProject(uint(idProject), uint(idUser))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
		}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Users removed from project successfully",
	})
}

func (h *ProjectHandler) AddTaskToProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	idTask, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}

	err = h.projectService.AddTaskToProject(uint(idProject), uint(idTask))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks added to project successfully",
	})
}

func (h *ProjectHandler) RemoveTaskFromProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	idTask, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}

	err = h.projectService.RemoveTaskFromProject(uint(idProject), uint(idTask))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks removed from project successfully",
	})
}