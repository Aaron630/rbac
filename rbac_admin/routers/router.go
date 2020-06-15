package routers

import (
	"rbac_admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/auth/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/auth/logout", &controllers.AuthController{}, "post:LoginOut")

	beego.Router("/user/info", &controllers.UserController{}, "get:UserInfo")

}
