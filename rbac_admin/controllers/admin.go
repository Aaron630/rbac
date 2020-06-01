package controllers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"rbac_admin/models"
	"strings"
	"time"

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
	type CustomeClaims struct {
		user string
		iat  int64
		jwt.StandardClaims
	}
	customeClaims := &CustomeClaims{
		user: adminUser.Account,
		iat:  time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customeClaims)
	tokenString, err := token.SignedString([]byte("ppp"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenString)
	// json webtoken 解析
	tt, err := jwt.ParseWithClaims(tokenString, &CustomeClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		// fmt.Println(((*token).Claims))
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("ppp"), nil
	})
	b := make([]byte, 2)
	rand.Read(b)
	fmt.Println(hex.EncodeToString(passwdMD5[:]))
	fmt.Println("aaaaaa:", tt)
}
