package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lucapierini/project-go-task_manager/models"
	"gorm.io/gorm"
)

const (
    accessTokenDuration  = 15 * time.Minute
    refreshTokenDuration = 7 * 24 * time.Hour
)

var (
    jwtSecret        = []byte(os.Getenv("JWT_SECRET"))
    ErrInvalidToken  = errors.New("invalid token")
    ErrExpiredToken  = errors.New("token has expired")
    ErrUnauthorized  = errors.New("unauthorized")
)

type Claims struct {
    UserID uint
    Roles  []string
    TokenType string
    jwt.StandardClaims
}

type TokenPair struct {
    AccessToken  string
    RefreshToken string
}

func GenerateTokenPair(user *models.User) (*TokenPair, error) {
    // Generate access token
    accessToken, err := generateToken(user, "access", accessTokenDuration)
    if err != nil {
        return nil, fmt.Errorf("error generating access token: %w", err)
    }

    // Generate refresh token
    refreshToken, err := generateToken(user, "refresh", refreshTokenDuration)
    if err != nil {
        return nil, fmt.Errorf("error generating refresh token: %w", err)
    }

    return &TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

func generateToken(user *models.User, tokenType string, duration time.Duration) (string, error) {
    var roleNames []string
    for _, role := range user.Roles {
        roleNames = append(roleNames, role.Name)
    }

    claims := Claims{
        UserID: user.ID,
        Roles:  roleNames,
        TokenType: tokenType,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(duration).Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtSecret, nil
    })

    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorExpired != 0 {
                return nil, ErrExpiredToken
            }
        }
        return nil, ErrInvalidToken
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, ErrInvalidToken
    }

    return claims, nil
}

func RefreshTokenHandler(c *gin.Context) {
    refreshToken := c.GetHeader("Refresh-Token")
    if refreshToken == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token required"})
        return
    }

    claims, err := ValidateToken(refreshToken)
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
    }

    // Generate new token pair
    tokenPair, err := GenerateTokenPair(user)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to generate new tokens"})
        return
    }

    c.JSON(http.StatusOK, tokenPair)
}

func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := ValidateToken(tokenString)
        if err != nil {
            if err == ErrExpiredToken {
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

        c.Set("userID", claims.UserID)
        c.Set("userRoles", claims.Roles)
        c.Next()
    }
}