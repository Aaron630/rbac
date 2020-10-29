package main

import (
	"fmt"
	"rbaca/middleawares"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleawares.AuthPermissionMiddleaware())
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("999999")
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
