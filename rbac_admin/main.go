package main

import (
	_ "rbac/models"

	"rbac/middleawares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleawares.Consuming)
	r.Use(middleawares.Auth)
	Router(r)
	r.Run(":8080")
}
