package main

import (
	_ "rbac/models"

	"rbac/middleawares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	store := cookie.NewStore([]byte("OWXMEWZTKK"))
	r.Use(sessions.Sessions("gsession", store))

	r.Use(middleawares.Consuming)
	r.Use(middleawares.Auth)
	Router(r)
	r.Run(":8080")
}
