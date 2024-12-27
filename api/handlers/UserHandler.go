package handlers

import (
	"github.com/gin-gonic/gin"
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

func (h *UserHandler) RegisterUser(c *gin.Context){
	users, err := 
}