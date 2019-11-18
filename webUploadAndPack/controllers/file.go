package controllers

import (
	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "file.tpl"
}
