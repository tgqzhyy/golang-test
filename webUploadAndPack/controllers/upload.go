package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "upload.tpl"
}

func (c *UploadController) Post() {
	c.TplName = "upload.tpl"
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("getFile err:", err)
		c.Ctx.Redirect(302, "/erro")
		return
	} else {
		fileName := h.Filename
		ExeFile := strings.Split(fileName, ".")
		layout := strings.ToLower(ExeFile[len(ExeFile)-1])
		if layout != "exe" && layout != "dll" {
			c.Ctx.WriteString("请上传符合格式的内容(exe、dll)")
			return
		}
		fileNamePack := beego.AppConfig.String("PackDiskPath") + strconv.FormatInt(time.Now().UnixNano(), 10) + h.Filename
		err = c.SaveToFile("file", fileNamePack) //Save disk path
		if err != nil {
			//c.Ctx.WriteString("File upload failed!")
			c.Ctx.Redirect(302, "/erro")
		} else {
			//c.Ctx.WriteString("File upload succed!")
			c.Ctx.Redirect(302, "/file")
			go AddPack(fileNamePack)
		}
	}
	defer f.Close()

}

func AddPack(f string) {
	cmd := exec.Command("ls", "-a", "-l", ">>", "bb.txt")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		log.Fatal(f)
	}

	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(opBytes))
}
