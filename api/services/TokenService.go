package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lucapierini/project-go-task_manager/models"
)


var (
    jwtSecret        = []byte(os.Getenv("JWT_SECRET"))
    ErrInvalidToken  = errors.New("invalid token")
    ErrExpiredToken  = errors.New("token has expired")
    ErrUnauthorized  = errors.New("unauthorized")
)

const (
    accessTokenDuration  = 15 * time.Minute
    refreshTokenDuration = 7 * 24 * time.Hour
)

func ValidateToken(tokenString string) (*models.Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
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

    claims, ok := token.Claims.(*models.Claims)
    if !ok || !token.Valid {
        return nil, ErrInvalidToken
    }

    return claims, nil
}

func GenerateTokenPair(user *models.User) (*models.TokenPair, error) {
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

    return &models.TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

func generateToken(user *models.User, tokenType string, duration time.Duration) (string, error) {
    var roleNames []string
    for _, role := range user.Roles {
        roleNames = append(roleNames, role.Name)
    }

    claims := models.Claims{
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
