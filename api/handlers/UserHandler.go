package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/middlewares"
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

func (h *UserHandler) Register(c *gin.Context) {
    var userDto dto.UserDto
    if err := c.ShouldBindJSON(&userDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request data",
            "details": err.Error(),
        })
        return
    }

    user, err := h.userService.RegisterUser(userDto)
    if err != nil {
        statusCode := http.StatusInternalServerError
        if err == services.ErrEmailAlreadyRegistered {
            statusCode = http.StatusConflict
        } else if err == services.ErrInvalidData {
            statusCode = http.StatusBadRequest
        }

        c.JSON(statusCode, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context){
	var loginDto dto.LoginDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	fmt.Println(loginDto)

	user, err := h.userService.LoginUser(loginDto)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("generating token for: " + user.Username)
	// token, err := utils.GenerateToken(user)
	token, err := middlewares.GenerateToken(user)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("auth_token", token, int(time.Now().Add(1*time.Hour).Unix()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	// Responder con el token
	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("Authorization", token, 3600 * 24 * 30 ,"", "", false, true)

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