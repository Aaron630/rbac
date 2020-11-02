package main

import (
	"fmt"
	. "rbac/controllers"
	. "rbac/middleawares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Router func
func Router(r *gin.Engine) {
	r.GET("/ping", VerifyPermission("a"), func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println(session.Get("id"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRouter := r.Group("/auth")
	authController := &AuthController{}
	{
		authRouter.POST("/login", authController.Login)
		authRouter.GET("/logout", authController.LoginOut)
	}
}
