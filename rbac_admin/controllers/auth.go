package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"rbac_admin/models"
	"rbac_admin/utils"
	"strings"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

type login struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

func (c *AuthController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
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
		// TODO
	}
	passwdMD5 := md5.Sum([]byte(loginInfo.Password))
	passwdStr := strings.ToUpper(hex.EncodeToString(passwdMD5[:]))
	fmt.Println(passwdStr)

	if passwdStr != adminUser.Passwd {
		// 密码错误
	}
	accessToken, err := utils.GenerateAssesToken(adminUser.Account)
	if err != nil {
		// TODO
	}
	fmt.Println(accessToken)
	data := map[string]interface{}{
		"userId":      adminUser.ID,
		"username":    adminUser.Account,
		"avatar":      adminUser.Avatar,
		"roleId":      "admin", // TODO
		"accessToken": accessToken,
	}
	c.Data["json"] = &data
	c.ServeJSON()
}

func (c *AuthController) LoginOut() {
	data := map[string]interface{}{}
	c.Data["json"] = &data
	c.ServeJSON()
}
