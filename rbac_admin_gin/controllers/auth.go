package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"rbac/models"
	"strings"

	// "rbac/utils"

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
	// fmt.Println(passwdStr)

	if passwdStr != adminUser.Passwd {
		c.JSON(http.StatusBadRequest, gin.H{"message": "密码错误"})
		return
	}
	// // accessToken, err := utils.GenerateAssesToken(adminUser.Account)
	// if err != nil {
	// 	c.Error(1001, "系统错误")
	// }
	// // fmt.Println(accessToken)
	// data := map[string]interface{}{
	// 	"userId":   adminUser.ID,
	// 	"username": adminUser.Account,
	// 	"avatar":   adminUser.Avatar,
	// 	"roleId":   "admin", // TODO
	// 	// "accessToken": accessToken,
	// }
	c.JSON(200, gin.H{
		"sussces": "ok",
	})
}

// LoginOut func
func (p *AuthController) LoginOut(c *gin.Context) {
	// data := map[string]interface{}{}
	// c.Success(data, "登出成功")
}
