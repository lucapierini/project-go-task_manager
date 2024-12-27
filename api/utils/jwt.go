package utils

import 
(
	"os"
	"github.com/golang-jwt/jwt"
	"github.com/lucapierini/project-go-task_manager/models"
	"time"
)


var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(user models.User) (string,error){
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

func ValidateToken(tokenString string) (jwt.MapClaims,error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil,err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims,nil
	} else {
		return nil,err
	}
}