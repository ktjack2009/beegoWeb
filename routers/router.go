package routers

import (
	"beegoWeb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由
	beego.Router("/", &controllers.MainController{})

	// 正则路由
	beego.Router("/user/:username([\\w]+)", &controllers.UserController{})
	beego.Router("/api/:id([0-9]+)", &controllers.ApiController{})

	// http://localhost:8080/download/123123/123.json   ext:json path:123123/123
	beego.Router("/download/*.*", &controllers.ApiController{})
	beego.Router("/cms_:id([0-9]+).html", &controllers.ApiController{})
}
