package middleawares

import (
	"fmt"
	"rbac/models"
	"rbac/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// VerifyPermission func
func VerifyPermission(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("id")
		url := c.Request.URL.String()
		adminUser := &models.Admin{}
		_, modules := adminUser.GetUserPermissions(userID.(uint32))
		isPass := utils.Contain(name, modules, "Action")
		if isPass == false {
			c.JSON(200, gin.H{
				"sussces": "no",
				"msg":     "权限不够",
			})
			c.Abort()
		}
		fmt.Printf("[permission]the request URL: %s permissionName %s\n", url, name)
	}
}
