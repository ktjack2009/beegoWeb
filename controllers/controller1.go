package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	username := c.Ctx.Input.Param(":username")
	c.Ctx.WriteString("user:" + username)
}

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	parmas := c.Ctx.Input.Params()
	res := ""
	for k, v := range parmas {
		res = res + k[1:] + ":" + v + "\n"
	}
	c.Ctx.WriteString(res)
}
