package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/models"
	"github.com/lucapierini/project-go-task_manager/services"
	"gorm.io/gorm"
)

func RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.GetHeader("Refresh-Token")
	if refreshToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token required"})
		return
	}

	claims, err := services.ValidateToken(refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims.TokenType != "refresh" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token type"})
		return
	}

	// Create a minimal user object with the claims data
	user := &models.User{
		Model: gorm.Model{ID: claims.UserID},
		Roles: []models.Role{},
	}

	for _, roleName := range claims.Roles {
		user.Roles = append(user.Roles, models.Role{Name: roleName})
	}

	// Generate new token pair
	tokenPair, err := services.GenerateTokenPair(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to generate new tokens"})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}