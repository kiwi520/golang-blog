package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "roczhmg@gmail.com"
	c.TplName = "index.tpl"
}

func (this *UserController) Get() {
	this.TplName = "layout/layout.html"
	//this.Layout = "layout/layout.html"
}