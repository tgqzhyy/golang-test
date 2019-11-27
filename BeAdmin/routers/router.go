package routers

import (
	"golang-test/BeAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	apins := beego.NewNamespace("/api",
		beego.NSNamespace("v1",
			beego.NSNamespace("/update",
				beego.NSRouter("/postsql",&controllers.PostSqlController{}),
				),
		),)
	beego.AddNamespace(apins)
    beego.Router("/", &controllers.MainController{})
}
