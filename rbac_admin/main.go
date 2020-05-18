package main

import (
	_ "rbac/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

