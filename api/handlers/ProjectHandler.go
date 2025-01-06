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
	id, err := strconv.Atoi(c.Param("id"))

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
	id, err := strconv.Atoi(c.Param("id"))

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
	id, err := strconv.Atoi(c.Param("id"))

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
	id, err := strconv.Atoi(c.Param("id"))

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

func (h *ProjectHandler) AddUsersToProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
		return
	}

	udsString := c.QueryArray("users_ids")
	var usersIds []uint
	for _, id := range udsString {
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
			return
			}
		usersIds = append(usersIds, uint(userId))
	}
	err = h.projectService.AddUsersToProject(uint(idProject), usersIds)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
		}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Users added to project successfully",
	})
}

func (h *ProjectHandler) RemoveUsersFromProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
		return
	}

	udsString := c.QueryArray("users_ids")
	var usersIds []uint
	for _, id := range udsString {
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
			return
			}
		usersIds = append(usersIds, uint(userId))
	}

	err = h.projectService.RemoveUsersFromProject(uint(idProject), usersIds)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
		}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Users removed from project successfully",
	})
}

func (h *ProjectHandler) AddTasksToProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("id_project"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	tasksString := c.QueryArray("tasks_ids")
	var tasksIds []uint
	for _, id := range tasksString {
		taskId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
			return
		}
		tasksIds = append(tasksIds, uint(taskId))
	}

	idUser, err := strconv.Atoi(c.Param("id_user"))

	err = h.projectService.AddTasksToProject(uint(idProject), tasksIds, uint(idUser))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks added to project successfully",
	})
}

func (h *ProjectHandler) RemoveTasksFromProject(c *gin.Context){
	idProject, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	tasksString := c.QueryArray("tasks_ids")
	var tasksIds []uint
	for _, id := range tasksString {
		taskId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project or user id"})
			return
		}
		tasksIds = append(tasksIds, uint(taskId))
	}

	err = h.projectService.RemoveTasksFromProject(uint(idProject), tasksIds)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks removed from project successfully",
	})
}