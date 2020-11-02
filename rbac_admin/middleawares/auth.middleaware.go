package middleawares

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Auth func
func Auth(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	url := c.Request.URL.String()
	if id == nil && url != "/auth/login" && url != "/ping" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请先登入"})
		fmt.Printf("[unlogin]the request URL %s %d\n", url, id)
		return
	}
}
