package main

import (
	. "rbac/controllers"
	. "rbac/middleawares"

	"github.com/gin-gonic/gin"
)

// Router func
func Router(r *gin.Engine) {
	r.GET("/ping", VerifyPermission("ping"), func(c *gin.Context) {
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

	userRouter := r.Group("/user")
	userController := &UserController{}
	{
		userRouter.GET("/userInfo", userController.UserInfo)
	}
}
