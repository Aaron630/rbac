package controllers

import (
	"rbac/models"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct {
	BaseController
}

// UserInfo func
func (p *UserController) UserInfo(c *gin.Context) {
	adminUser := &models.Admin{}
	session := sessions.Default(c)
	var userID = session.Get("id")
	adminRoles, modules := adminUser.GetUserPermissions(userID.(uint32))
	c.JSON(200, gin.H{
		"sussces": "ok",
		"data": map[string]interface{}{
			"userId":   adminUser.ID,
			"username": adminUser.Account,
			"avatar":   adminUser.Avatar,
			"roles":    adminRoles,
			"modules":  modules,
		},
	})
}
