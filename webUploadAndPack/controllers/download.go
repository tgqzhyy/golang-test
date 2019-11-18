package controllers

import (
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "file.tpl"
}
