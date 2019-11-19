package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"path"
)

type FileController struct {
	beego.Controller
}
type Content struct {
	Name string				//文件名
	Size int64				//大小
	Extension string 		//类型/扩展名
	IsDir bool				//是否是目录
	LastModified string		//修改日期
	Path string				//路径
	//Md5 string`json:"md5"`
}

// 获取文件夹列表
func (c *FileController)getDirContent(p string,s bool)([]Content,error)  {
	files,_:=ioutil.ReadDir(path.Join(beego.AppConfig.String("SavaDiskPath"),p))

	var contents []Content

	for _,f :=range files{

		content :=Content{
			Name:         f.Name(),
			Size:         f.Size()/1024/1024,
			IsDir:        f.IsDir(),
			LastModified: f.ModTime().String(),
			Path: path.Join(p,f.Name()), //fuck~~~
		}
		if !s{
			if !f.IsDir(){
				content.Extension = path.Ext(p + f.Name())
				contents =append(contents,content)
			}
		}else {
			if f.IsDir(){
				contents =append(contents,content)
			}
		}
	}
	return contents,nil
}

func (c *FileController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "file.tpl"
	//beego.Notice("打开file")

	//beego.Notice(c.GetString("dir"))
	if cc:=c.GetString("dir");cc==""{
		ct,_ :=c.getDirContent("",true)
		ctt,_ :=c.getDirContent("",false)
		c.Data["filelist"]=ct
		c.Data["contentlist"]=ctt
	}else {
		ct,_ :=c.getDirContent(cc,true)
		ctt,_ :=c.getDirContent(cc,false)
		c.Data["filelist"]=ct
		c.Data["contentlist"]=ctt
	}
}
