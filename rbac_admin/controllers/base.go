package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (t *BaseController) Success(data interface{}, message string) {
	ret := map[string]interface{}{"code": 1000, "message": message}
	ret["result"] = data
	t.Data["json"] = ret
	t.ServeJSON()
	t.StopRun()
}

func (t *BaseController) Error(code int16, message string) {
	ret := map[string]interface{}{"code": code, "message": message}
	t.Data["json"] = ret
	t.ServeJSON()
	t.StopRun()
}
