package utils

import (
	"github.com/lucapierini/project-go-task_manager/responses"
	"github.com/gin-gonic/gin"
)

func GetUserInfoFromContext(c *gin.Context) *responses.UserInfo {
	userInfo, _ := c.Get("UserInfo")

	user, _ := userInfo.(*responses.UserInfo)

	return user
}
