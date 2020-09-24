package controllers

import (
	"github.com/astaxie/beego"
)

type PostSqlController struct {
	beego.Controller
}

func (c *PostSqlController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
