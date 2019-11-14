package routers

import (
	"golang-test/毒鸡汤/djt/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
