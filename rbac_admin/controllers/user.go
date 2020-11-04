package controllers

import (
	"fmt"
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
	adminUser.GetUserPermissions(userID.(uint32))
	fmt.Println(adminUser)
}
