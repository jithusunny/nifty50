package routers

import (
	"github.com/astaxie/beego"
	"github.com/jithusunny/nifty50/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
