package routers

import (
	"beegoWeb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}) // 将路由映射到controller
	beego.Router("/index", &controllers.IndexController{})
}
