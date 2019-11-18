package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"path"
)

type FileController struct {
	beego.Controller
}
type Content struct {
	Name string`json:"name"`						//文件名
	Size int64`json:"size"`							//大小
	Extension string `json:"extension,omitempty"` 	//类型/扩展名
	IsDir bool`josn:"is_dir"`						//是否是目录
	LastModified string`json:"lastModified"`		//修改日期
	//Md5 string`json:"md5"`
}

// 获取文件夹列表
func (c *FileController)getDirContent(p string)([]byte,error)  {
	files,_:=ioutil.ReadDir(p)

	var contents []Content

	for _,f :=range files{
		content :=Content{
			Name:         f.Name(),
			Size:         f.Size()/1024,
			IsDir:        f.IsDir(),
			LastModified: f.ModTime().String(),
		}
		if !f.IsDir(){
			content.Extension = path.Ext(p + f.Name())
		}
		contents =append(contents,content)

	}
	//beego.Notice("fuck")
	//beego.Notice(contents)
	b,err := json.Marshal(contents)
	if err !=nil{
		return nil,err
	}
	return b,nil
}


func (c *FileController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "file.tpl"
	beego.Notice("打开file")
	ct,_ :=c.getDirContent(beego.AppConfig.String("SavaDiskPath"))
	c.Data["ct"]=ct

	beego.Notice(string(ct))
}
