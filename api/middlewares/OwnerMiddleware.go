// OwnerMiddleware.go
package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/models"
)

func IsOwner(resourceType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el usuario del contexto (establecido por AuthMiddleware)
		userClaims, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
			return
		}
		claims := userClaims.(*models.Claims)

		// Verificar si el usuario es administrador
		isAdmin := false
		for _, role := range claims.Roles {
			if role == "Administrador" {
				isAdmin = true
				break
			}
		}

		// Si es administrador, permitir acceso
		if isAdmin {
			c.Next()
			return
		}

		// Verificar propiedad según el tipo de recurso
		var isOwner bool

		switch resourceType {
		case "user":
			// Obtener el ID del recurso de los parámetros
			// resourceID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
            resourceID, err := strconv.Atoi(c.Param("userId"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid resource ID"})
				return
			}
			isOwner = claims.UserID == uint(resourceID)

		case "project":
			// Obtener el ID del recurso de los parámetros
			resourceID, err := strconv.ParseUint(c.Param("projectId"), 10, 64)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid resource ID"})
				return
			}
			var project models.Project
			result := config.DB.First(&project, resourceID)
			if result.Error != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "project not found"})
				return
			}
			isOwner = project.OwnerID == claims.UserID

		case "task":
			// Obtener el ID del recurso de los parámetros
			resourceID, err := strconv.ParseUint(c.Param("taskId"), 10, 64)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid resource ID"})
				return
			}
			var task models.Task
			result := config.DB.First(&task, resourceID)
			if result.Error != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "task not found"})
				return
			}
			isOwner = task.OwnerID == claims.UserID

		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid resource type"})
			return
		}

		if !isOwner {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you don't have permission to modify this resource"})
			return
		}

		c.Next()
	}
}
