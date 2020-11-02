package middleawares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// VerifyPermission func
func VerifyPermission(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		fmt.Printf("[permission]the request URL: %s permissionName %s\n", url, name)
	}
}
