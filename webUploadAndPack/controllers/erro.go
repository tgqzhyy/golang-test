package controllers

import (
	"github.com/astaxie/beego"
)

type ErroController struct {
	beego.Controller
}

func (c *ErroController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "erro.tpl"
}
