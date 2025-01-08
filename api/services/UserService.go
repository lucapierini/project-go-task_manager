package services

import (
	"errors"
	"fmt"

	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/models"
	// "github.com/lucapierini/project-go-task_manager/utils"
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
	AssignRoleToUser(userId uint, roleId uint) error
	UnassignRoleToUser(userId uint, roleId uint) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrInvalidData            = errors.New("invalid data provided")
)

func (s *UserService) RegisterUser(userDto dto.UserDto) (*models.User, error) {
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
	} else {
		// If no roles are specified, assign default role
		var defaultRole models.Role
		if err := config.DB.First(&defaultRole, "name = ?", "Usuario").Error; err != nil {
			return nil, err
		}
		user.Roles = []models.Role{defaultRole}
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}


func (s *UserService) LoginUser(loginDto dto.LoginDto) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Roles").Where("email = ?", loginDto.Email).First(&user).Error; err != nil {
		fmt.Println("invalid credentials")
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		fmt.Println("invalid password")
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}


func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Roles").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
    var user models.User
    if err := config.DB.Preload("Roles").Where("id = ?", id).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UserService) UpdateUser(id uint, userDto dto.UserDto) (*models.User, error) {
	user, err := s.GetUserById(id)
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



func (s *UserService) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) DeleteUser (id uint) error {
    if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
        return err
    }
    return nil
}

func (s *UserService) AssignRoleToUser(userId uint, roleId uint) error {
	var user models.User
	if err := config.DB.Preload("Roles").Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

		// Check if role is already assigned to the user
	for _, r := range user.Roles {
		if r.ID == roleId {
			return errors.New("role already assigned to user")
		}
	}

	var role models.Role
	if err := config.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		return err
	}

	err := config.DB.Model(&user).Association("Roles").Append(&role)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UnassignRoleToUser(userId uint, roleId uint) error {
	var user models.User
	if err := config.DB.Preload("Roles").Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

		// Check if role is already assigned to the user
	var found bool
	for i, r := range user.Roles {
		if r.ID == roleId {
			user.Roles = append(user.Roles[:i], user.Roles[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("role not assigned to user")
	}

	var role models.Role
	if err := config.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		return err
	}

	err := config.DB.Model(&user).Association("Roles").Delete(&role)

	if err != nil {
		return err
	}

	return nil
}