package main

import (
	. "rbaca/controllers"
	. "rbaca/middleawares"

	"github.com/gin-gonic/gin"
)

// Router func
func Router(r *gin.Engine) {
	r.GET("/ping", VerifyPermission("a"), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRouter := r.Group("/auth")
	authController := &AuthController{}
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/logout", authController.LoginOut)
	}
}
