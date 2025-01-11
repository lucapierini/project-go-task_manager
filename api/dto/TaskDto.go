package dto

type TaskDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerID     uint   `json:"owner_id" binding:"required"`
	ProjectID  uint   `json:"project_id"`
}