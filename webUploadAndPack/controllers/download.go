package controllers

import (
	"github.com/astaxie/beego"
	"path"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "file.tpl"
	d := path.Join(beego.AppConfig.String("SavaDiskPath"), c.GetString("d"))
	beego.Notice(d)
	c.Ctx.Output.Download(d)

}
