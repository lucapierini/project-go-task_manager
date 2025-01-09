package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/services"
)

func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := services.ValidateToken(tokenString)
        if err != nil {
            if err == services.ErrExpiredToken {
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired", "code": "TOKEN_EXPIRED"})
                return
            }
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        if claims.TokenType != "access" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token type"})
            return
        }

        if len(requiredRoles) > 0 {
            hasRequiredRole := false
            for _, requiredRole := range requiredRoles {
                for _, userRole := range claims.Roles {
                    if requiredRole == userRole {
                        hasRequiredRole = true
                        break
                    }
                }
                if hasRequiredRole {
                    break
                }
            }

            if !hasRequiredRole {
                c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
                return
            }
        }
        c.Set("user", claims)
        c.Next()
    }
}