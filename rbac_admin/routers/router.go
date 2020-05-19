package routers

import (
	"rbac_admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
