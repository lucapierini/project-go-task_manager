package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/services"
)

type RoleHandler struct {
	roleService services.RoleInterface
}

func NewRoleHandler(roleService services.RoleInterface) *RoleHandler {
    return &RoleHandler{
        roleService: roleService,
    }
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var roleDto dto.RoleDto
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	role, err := h.roleService.CreateRole(roleDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, role)
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	role, err := h.roleService.GetRoleById(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(200, role)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var roleDto dto.RoleDto
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	role, err := h.roleService.UpdateRole(uint(id), roleDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, role)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.roleService.DeleteRole(uint(id)); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Role deleted successfully"})
}

func (h *RoleHandler) ListRoles(c *gin.Context) {
	roles, err := h.roleService.ListRoles()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch roles"})
		return
	}

	c.JSON(200, roles)
}