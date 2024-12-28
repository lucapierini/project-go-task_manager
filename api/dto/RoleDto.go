package dto

type RoleDto struct {
	Name string `json:"name" binding:"required"`
}