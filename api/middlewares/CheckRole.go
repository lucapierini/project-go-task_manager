package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// RoleMiddleware verifica si el usuario tiene uno de los roles permitidos
func RoleMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el rol del usuario desde el contexto o el token
		userRole := c.GetHeader("Role") // Suponiendo que el rol se pasa en el header

		// Verificar si el rol del usuario está en la lista de roles permitidos
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next() // El rol es permitido, continuar con la siguiente función
				return
			}
		}

		// Si el rol no es permitido, devolver un error 403 Forbidden
		c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado"})
		c.Abort() // Abortamos la ejecución de la siguiente función
	}
}