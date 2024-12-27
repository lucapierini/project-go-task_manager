package dto

type UserDto struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    RoleIds  []uint `json:"role_ids"`
}