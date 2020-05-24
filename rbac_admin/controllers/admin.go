package controllers

import (
	"rbac_admin/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

// TODO
func (c *AdminController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	adminUser := &models.AdminUser{}
	adminUser.AdminUserInfo()
	// fmt.Println(tt)
}

func (c *AdminController) Post() {
	// adminUser := &models.AdminUser{}
	// tt := adminUser.AdminUserInfo()
	// fmt.Println(tt)
}
