package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller // controller实现beego.Controller自带的所有方法
}

func (c *MainController) Get() { // 改写Get方法
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl" // 指定渲染的模版
}
