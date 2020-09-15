package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"rbac/models"
	"rbac/utils"
	"strings"
)

type AuthController struct {
	BaseController
}

type login struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

func (c *AuthController) Login() {
	loginInfo := &login{}
	json.Unmarshal(c.Ctx.Input.RequestBody, loginInfo)
	// if err := c.ParseForm(loginInfo); err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(loginInfo)

	adminUser := &models.AdminUser{}
	adminUser.GetUserInfoByAccount(loginInfo.Account)
	if adminUser == nil {
		c.Error(1002, "账号不存在")
	}
	passwdMD5 := md5.Sum([]byte(loginInfo.Password))
	passwdStr := strings.ToUpper(hex.EncodeToString(passwdMD5[:]))
	fmt.Println(passwdStr)

	if passwdStr != adminUser.Passwd {
		c.Error(1003, "密码错误")
	}
	accessToken, err := utils.GenerateAssesToken(adminUser.Account)
	if err != nil {
		c.Error(1001, "系统错误")
	}
	fmt.Println(accessToken)
	data := map[string]interface{}{
		"userId":      adminUser.ID,
		"username":    adminUser.Account,
		"avatar":      adminUser.Avatar,
		"roleId":      "admin", // TODO
		"accessToken": accessToken,
	}
	c.Success(data, "登入成功")
}

func (c *AuthController) LoginOut() {
	data := map[string]interface{}{}
	c.Success(data, "登出成功")
}
