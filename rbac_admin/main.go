package main

import (
	_ "rbac_admin/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
