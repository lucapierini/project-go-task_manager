package services

import (
	"errors"

	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/models"
	"github.com/lucapierini/project-go-task_manager/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	RegisterUser(userDto dto.UserDto) (*models.User, error)
	LoginUser(loginDto dto.LoginDto) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ListUsers() ([]models.User, error)
	UpdateUser(id uint, userDto dto.UserDto) (*models.User, error)
	DeleteUser(id uint) error
}

func RegisterUser(userDto dto.UserDto) (*models.User, error) {
	// Check if email already exists
	var existingUser models.User
	if result := config.DB.Where("email = ?", userDto.Email).First(&existingUser); result.Error == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Password: string(hashedPassword),
	}

	// Add roles if specified
	if len(userDto.RoleIds) > 0 {
		var roles []models.Role
		if err := config.DB.Find(&roles, userDto.RoleIds).Error; err != nil {
			return nil, err
		}
		user.Roles = roles
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Roles").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id uint) (*models.User, error) {
    var user models.User
    if err := config.DB.Preload("Roles").Where("id = ?", id).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func UpdateUser(id uint, userDto dto.UserDto) (*models.User, error) {
	user, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	user.Username = userDto.Username
	user.Email = userDto.Email

	if userDto.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	// Update roles if specified
	if len(userDto.RoleIds) > 0 {
		var roles []models.Role
		if err := config.DB.Find(&roles, userDto.RoleIds).Error; err != nil {
			return nil, err
		}
		user.Roles = roles
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id uint) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func ListUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func LoginUser(loginDto dto.LoginDto) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", loginDto.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
