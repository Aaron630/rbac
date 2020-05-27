package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"rbac_admin/models"
	"strings"
	"time"

	"crypto/md5"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
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
	adminUser.GetUserInfoByAccount(loginInfo.Account)
	if adminUser == nil {
		// TODO
	}
	passwdMD5 := md5.Sum([]byte(loginInfo.Password))
	passwdStr := strings.ToUpper(hex.EncodeToString(passwdMD5[:]))

	if passwdStr != adminUser.Passwd {
		// 密码错误
	}
	// 颁发一个有限期72小时的证书
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": adminUser.Account,
		"exp":  time.Now().Add(time.Hour * time.Duration(72)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte("ppp"))
	if err != nil {
		fmt.Println(err)
	}
	// json webtoken 解析

	fmt.Println(tokenString)
}
