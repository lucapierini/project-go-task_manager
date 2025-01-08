package dto

import (

)

type ProjectDto struct {
	Name string `json:"name" binding:"required"`
	Budget uint `json:"budget" binding:"required"`
	OwnerID uint `json:"owner_id" binding:"required"`
	UsersIds []uint `json:"users_ids"`
	TasksIds []uint `json:"tasks_ids"`
}