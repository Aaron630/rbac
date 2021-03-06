package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"rbac/models"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthController struct
type AuthController struct {
	BaseController
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login func
func (p *AuthController) Login(c *gin.Context) {
	loginInfo := &login{}
	err := c.ShouldBindJSON(loginInfo)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Error(err)
		return
	}
	fmt.Printf("login-------------: %s %s", loginInfo.Username, loginInfo.Password)

	adminUser := &models.Admin{}
	adminUser.GetUserInfoByAccount(loginInfo.Username)
	if adminUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "账号不存在"})
		return
	}
	passwdMD5 := md5.Sum([]byte(loginInfo.Password))
	passwdStr := strings.ToUpper(hex.EncodeToString(passwdMD5[:]))

	if passwdStr != adminUser.Passwd {
		c.JSON(http.StatusBadRequest, gin.H{"message": "密码错误"})
		return
	}
	session := sessions.Default(c)
	// set session
	session.Set("id", adminUser.ID)
	session.Set("username", adminUser.Account)
	session.Save()

	c.JSON(200, gin.H{
		"sussces": "ok",
		"data": map[string]interface{}{
			"userId":   adminUser.ID,
			"username": adminUser.Account,
			"avatar":   adminUser.Avatar,
		},
	})
}

// LoginOut func
func (p *AuthController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{
		"sussces": "登出成功",
	})
}
