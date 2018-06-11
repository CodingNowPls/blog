package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.striverfeng.cn"
	c.Data["Email"] = "striverfeng@gmail.com"
	c.TplName = "index.html"
}
