package routers

import (
	"github.com/astaxie/beego"
	"golang-test/毒鸡汤/djt/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
