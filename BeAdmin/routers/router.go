package routers

import (
	"github.com/astaxie/beego"
	"golang-test/BeAdmin/controllers"
)

func init() {
	apins := beego.NewNamespace("/api",
		beego.NSNamespace("v1",
			beego.NSNamespace("/update",
				beego.NSRouter("/postsql", &controllers.PostSqlController{}),
			),
		))
	beego.AddNamespace(apins)
	beego.Router("/", &controllers.MainController{})
}
