package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/services"
)

type UserHandler struct {
	userService services.UserInterface
}

func NewUserHandler(userService services.UserInterface) *UserHandler{
	return &UserHandler{
		userService : userService,
	}
}

func (h *UserHandler) Register(c *gin.Context){
	var userDto dto.UserDto

	if c.Bind(&userDto) != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	user, err := h.userService.RegisterUser(userDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) Login(c *gin.Context){
	var loginDto dto.LoginDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	token, err := h.userService.LoginUser(loginDto)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (h *UserHandler) GetUser(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Id"})
		return
	}

	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}

	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userService.ListUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(200, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var userDto dto.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	user, err := h.userService.UpdateUser(uint(id), userDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}