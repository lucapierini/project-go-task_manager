package handlers

import (

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