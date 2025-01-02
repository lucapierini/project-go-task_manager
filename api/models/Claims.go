package models

import (
		"github.com/golang-jwt/jwt"
)


type Claims struct {
    UserID uint
    Roles  []string
    TokenType string
    jwt.StandardClaims
}