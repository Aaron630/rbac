package controllers

import (
	"encoding/json"
	"fmt"
	"rbac_admin/models"

	"crypto/md5"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

type login struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

func (c *AdminController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *AdminController) Post() {
	loginInfo := &login{}
	json.Unmarshal(c.Ctx.Input.RequestBody, loginInfo)
	// if err := c.ParseForm(loginInfo); err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(loginInfo)
	adminUser := &models.AdminUser{}
	adminUser.AdminUserInfo(loginInfo.Account)
	if adminUser == nil {
		// TODO
	}
	passwdMD5 := md5.Sum([]byte(loginInfo.Password))
	fmt.Println(string(passwdMD5[:]))
	fmt.Println(adminUser)
}
